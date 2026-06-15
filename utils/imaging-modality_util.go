package utils

import (
	"strings"
	"time"
)

// =====================================================================
// DICOM wildcard helpers cho MWL C-FIND query (PS3.4 §C.2.2.2)
//
// Khi modality gửi C-FIND-RQ, các matching key có thể chứa:
//   - "" (rỗng) hoặc "*"       → universal matching, không filter
//   - "NGUYEN*"                 → prefix match
//   - "*HUNG*"                  → contains match
//   - "N?UYEN"                  → single-char wildcard
//   - "uid1\uid2"               → multi-value (chỉ UI, CS, AE)
//   - "20260520-20260528" (DA)  → date range
//
// Các helper bên dưới chuẩn hoá input từ gateway thành thứ HIS dùng được:
// pattern SQL LIKE, time.Time, multi-value slice…
// =====================================================================

// DicomLikePattern convert DICOM wildcard pattern sang SQL LIKE pattern.
//
//	DICOM "*" → SQL "%"
//	DICOM "?" → SQL "_"
//
// SQL special chars trong input ("%", "_", "\") được escape trước khi
// thay wildcard, nên kết quả an toàn (chống LIKE injection).
//
// Trả về:
//   - (pattern, true)  nếu có wildcard cần dùng LIKE (caller add WHERE clause)
//   - ("", false)      nếu input nil/rỗng/"*" → không filter (skip clause)
//
// Caller dùng kèm `ESCAPE '\'` trong SQL query để xử các ký tự escaped.
// Ví dụ Postgres:
//
//	if pattern, ok := utils.DicomLikePattern(req.PatientName); ok {
//	    q.Where("patient.name ILIKE ? ESCAPE '\\'", pattern)
//	}
func DicomLikePattern(input *string) (string, bool) {
	if input == nil {
		return "", false
	}
	s := strings.TrimSpace(*input)
	if s == "" || s == "*" {
		return "", false
	}
	// Escape SQL LIKE special chars trước (theo thứ tự: \ trước, sau đó % _).
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, "%", `\%`)
	s = strings.ReplaceAll(s, "_", `\_`)
	// Convert DICOM wildcards sang SQL wildcards.
	s = strings.ReplaceAll(s, "*", "%")
	s = strings.ReplaceAll(s, "?", "_")
	return s, true
}

// DicomDateRange parse DICOM DA value:
//
//	"YYYYMMDD"                  → from = to = date đó
//	"YYYYMMDD-YYYYMMDD"         → range full
//	"-YYYYMMDD"                 → open-start (chỉ to)
//	"YYYYMMDD-"                 → open-end (chỉ from)
//
// Trả về (from, to, ok). nil pointer nghĩa là không có ràng buộc đầu đó.
// ok=false khi input nil/rỗng → không filter.
//
// Lưu ý: trả về *time.Time tại 00:00:00 UTC. Nếu cần so sánh với DB column
// `timestamptz`, caller có thể truncate thêm phần giờ phù hợp.
func DicomDateRange(input *string) (from *time.Time, to *time.Time, ok bool) {
	if input == nil {
		return nil, nil, false
	}
	s := strings.TrimSpace(*input)
	if s == "" {
		return nil, nil, false
	}

	parts := strings.SplitN(s, "-", 2)

	// Single date "YYYYMMDD"
	if len(parts) == 1 {
		t, err := time.Parse("20060102", parts[0])
		if err != nil {
			return nil, nil, false
		}
		return &t, &t, true
	}

	// Range "from-to" với 1 đầu có thể rỗng
	if parts[0] != "" {
		if pt, err := time.Parse("20060102", parts[0]); err == nil {
			from = &pt
		}
	}
	if parts[1] != "" {
		if pt, err := time.Parse("20060102", parts[1]); err == nil {
			to = &pt
		}
	}
	if from == nil && to == nil {
		return nil, nil, false
	}
	return from, to, true
}

// DicomTimeRange parse DICOM TM value tương tự DicomDateRange nhưng cho time:
//
//	"HHMMSS"                    → from = to = time đó
//	"HHMMSS-HHMMSS"             → range
//	"-HHMMSS" / "HHMMSS-"       → open-ended
//
// Trả về string format "HH:MM:SS" để dễ build SQL TIME comparison.
// (Không return time.Time vì TIME không gắn date.)
func DicomTimeRange(input *string) (from string, to string, ok bool) {
	if input == nil {
		return "", "", false
	}
	s := strings.TrimSpace(*input)
	if s == "" {
		return "", "", false
	}

	parse := func(raw string) (string, bool) {
		if len(raw) < 4 {
			return "", false
		}
		// "HHMM" (4), "HHMMSS" (6), "HHMMSS.ffffff" (>6) — pad cho đủ HHMMSS
		hms := raw
		if dot := strings.IndexByte(hms, '.'); dot >= 0 {
			hms = hms[:dot]
		}
		switch len(hms) {
		case 4:
			hms += "00"
		case 6:
			// already OK
		default:
			return "", false
		}
		return hms[0:2] + ":" + hms[2:4] + ":" + hms[4:6], true
	}

	parts := strings.SplitN(s, "-", 2)
	if len(parts) == 1 {
		v, vok := parse(parts[0])
		if !vok {
			return "", "", false
		}
		return v, v, true
	}

	if parts[0] != "" {
		if v, vok := parse(parts[0]); vok {
			from = v
		}
	}
	if parts[1] != "" {
		if v, vok := parse(parts[1]); vok {
			to = v
		}
	}
	if from == "" && to == "" {
		return "", "", false
	}
	return from, to, true
}

// DicomMultiValue split input theo DICOM backslash separator "\".
// Áp dụng cho UI (multi-UID list), CS, AE (multi-code list).
//
//	"uid1\uid2\uid3"            → []string{"uid1", "uid2", "uid3"}
//	"CT\MR\DR"                  → []string{"CT", "MR", "DR"}
//	""                          → nil (caller skip filter)
//	"single"                    → []string{"single"}
//
// Trả về nil nếu input nil/rỗng (→ caller không add WHERE clause).
// Các giá trị empty sau split bị bỏ qua (vd "CT\\\\MR" → ["CT","MR"]).
func DicomMultiValue(input *string) []string {
	if input == nil {
		return nil
	}
	s := strings.TrimSpace(*input)
	if s == "" {
		return nil
	}
	raw := strings.Split(s, "\\")
	result := make([]string, 0, len(raw))
	for _, v := range raw {
		v = strings.TrimSpace(v)
		if v != "" {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// HasWildcard kiểm tra input có chứa DICOM wildcard ('*' hoặc '?') không.
// Hữu ích khi muốn quyết định dùng exact match (`=`) hay LIKE.
func HasWildcard(input *string) bool {
	if input == nil {
		return false
	}
	return strings.ContainsAny(*input, "*?")
}
