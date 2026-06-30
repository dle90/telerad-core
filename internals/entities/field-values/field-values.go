package fieldValues

const (
	// Giới tính bệnh nhân
	GENDER = "GENDER"

	// Telerad reading order — trạng thái ca đọc (cột status)
	TELERAD_READING_ORDER_STATUS = "TELERAD-READING-ORDER_STATUS"

	// Bộ phận chụp theo defined-term của PACS (DICOM Body Part Examined biến thể PACS).
	// Giữ NGUYÊN văn chuỗi PACS gửi (gồm cả các viết tắt / sai chính tả như FEMALEPELVIS,
	// MALEPELVIS) để khớp đúng khi lookup.
	PACS_BODY_PART = "PACS_BODY-PART"
)

var (
	// Giới tính
	GENDER_MALE   = ColumnValueString{Value: "MALE", Code: GENDER, Name: "Nam", ShortName: "Nam"}
	GENDER_FEMALE = ColumnValueString{Value: "FEMALE", Code: GENDER, Name: "Nữ", ShortName: "Nữ"}
)

var (
	// Telerad reading order status
	TELERAD_READING_ORDER_STATUS_UNREAD           = ColumnValueString{Value: "UNREAD", Code: TELERAD_READING_ORDER_STATUS, Name: "Chưa đọc", ShortName: "Chưa đọc"}                     // chưa đọc
	TELERAD_READING_ORDER_STATUS_READING          = ColumnValueString{Value: "READING", Code: TELERAD_READING_ORDER_STATUS, Name: "Đang đọc", ShortName: "Đang đọc"}                    // đang đọc
	TELERAD_READING_ORDER_STATUS_PENDING_APPROVAL = ColumnValueString{Value: "PENDING_APPROVAL", Code: TELERAD_READING_ORDER_STATUS, Name: "Đã đọc, chờ duyệt", ShortName: "Chờ duyệt"} // đã đọc xong, chờ duyệt
	TELERAD_READING_ORDER_STATUS_APPROVED         = ColumnValueString{Value: "APPROVED", Code: TELERAD_READING_ORDER_STATUS, Name: "Đã duyệt", ShortName: "Đã duyệt"}                   // đã duyệt
)

var (
	// PACS body part
	PACS_BODY_PART_BRAIN          = ColumnValueString{Value: "BRAIN", Code: PACS_BODY_PART, Name: "BRAIN", ShortName: "BRAIN"}
	PACS_BODY_PART_HEAD           = ColumnValueString{Value: "HEAD", Code: PACS_BODY_PART, Name: "HEAD", ShortName: "HEAD"}
	PACS_BODY_PART_LSPINE         = ColumnValueString{Value: "LSPINE", Code: PACS_BODY_PART, Name: "LSPINE", ShortName: "LSPINE"}
	PACS_BODY_PART_GENERALABDOMEN = ColumnValueString{Value: "GENERALABDOMEN", Code: PACS_BODY_PART, Name: "GENERALABDOMEN", ShortName: "GENERALABDOMEN"}
	PACS_BODY_PART_FEMALEPELVIS   = ColumnValueString{Value: "FEMALEPELVIS", Code: PACS_BODY_PART, Name: "FEMALEPELVIS", ShortName: "FEMALEPELVIS"}
	PACS_BODY_PART_FOREARM        = ColumnValueString{Value: "FOREARM", Code: PACS_BODY_PART, Name: "FOREARM", ShortName: "FOREARM"}
	PACS_BODY_PART_FOOT           = ColumnValueString{Value: "FOOT", Code: PACS_BODY_PART, Name: "FOOT", ShortName: "FOOT"}
	PACS_BODY_PART_FA             = ColumnValueString{Value: "FA", Code: PACS_BODY_PART, Name: "FA", ShortName: "FA"}
	PACS_BODY_PART_ST             = ColumnValueString{Value: "ST", Code: PACS_BODY_PART, Name: "ST", ShortName: "ST"}

	PACS_BODY_PART_ORBITS   = ColumnValueString{Value: "ORBITS", Code: PACS_BODY_PART, Name: "ORBITS", ShortName: "ORBITS"}
	PACS_BODY_PART_NECK     = ColumnValueString{Value: "NECK", Code: PACS_BODY_PART, Name: "NECK", ShortName: "NECK"}
	PACS_BODY_PART_CHEST    = ColumnValueString{Value: "CHEST", Code: PACS_BODY_PART, Name: "CHEST", ShortName: "CHEST"}
	PACS_BODY_PART_LIVER    = ColumnValueString{Value: "LIVER", Code: PACS_BODY_PART, Name: "LIVER", ShortName: "LIVER"}
	PACS_BODY_PART_SHOULDER = ColumnValueString{Value: "SHOULDER", Code: PACS_BODY_PART, Name: "SHOULDER", ShortName: "SHOULDER"}
	PACS_BODY_PART_WRIST    = ColumnValueString{Value: "WRIST", Code: PACS_BODY_PART, Name: "WRIST", ShortName: "WRIST"}
	PACS_BODY_PART_ANKLE    = ColumnValueString{Value: "ANKLE", Code: PACS_BODY_PART, Name: "ANKLE", ShortName: "ANKLE"}

	PACS_BODY_PART_SKULL         = ColumnValueString{Value: "SKULL", Code: PACS_BODY_PART, Name: "SKULL", ShortName: "SKULL"}
	PACS_BODY_PART_CSPINE        = ColumnValueString{Value: "CSPINE", Code: PACS_BODY_PART, Name: "CSPINE", ShortName: "CSPINE"}
	PACS_BODY_PART_BREAST        = ColumnValueString{Value: "BREAST", Code: PACS_BODY_PART, Name: "BREAST", ShortName: "BREAST"}
	PACS_BODY_PART_GENERALPELVIS = ColumnValueString{Value: "GENERALPELVIS", Code: PACS_BODY_PART, Name: "GENERALPELVIS", ShortName: "GENERALPELVIS"}
	PACS_BODY_PART_KNEE          = ColumnValueString{Value: "KNEE", Code: PACS_BODY_PART, Name: "KNEE", ShortName: "KNEE"}
	PACS_BODY_PART_LEG           = ColumnValueString{Value: "LEG", Code: PACS_BODY_PART, Name: "LEG", ShortName: "LEG"}
	PACS_BODY_PART_CB            = ColumnValueString{Value: "CB", Code: PACS_BODY_PART, Name: "CB", ShortName: "CB"}
	PACS_BODY_PART_TB            = ColumnValueString{Value: "TB", Code: PACS_BODY_PART, Name: "TB", ShortName: "TB"}

	PACS_BODY_PART_PITUITARY  = ColumnValueString{Value: "PITUITARY", Code: PACS_BODY_PART, Name: "PITUITARY", ShortName: "PITUITARY"}
	PACS_BODY_PART_TSPINE     = ColumnValueString{Value: "TSPINE", Code: PACS_BODY_PART, Name: "TSPINE", ShortName: "TSPINE"}
	PACS_BODY_PART_ABDOMEN    = ColumnValueString{Value: "ABDOMEN", Code: PACS_BODY_PART, Name: "ABDOMEN", ShortName: "ABDOMEN"}
	PACS_BODY_PART_MALEPELVIS = ColumnValueString{Value: "MALEPELVIS", Code: PACS_BODY_PART, Name: "MALEPELVIS", ShortName: "MALEPELVIS"}
	PACS_BODY_PART_HAND       = ColumnValueString{Value: "HAND", Code: PACS_BODY_PART, Name: "HAND", ShortName: "HAND"}
	PACS_BODY_PART_FEMUR      = ColumnValueString{Value: "FEMUR", Code: PACS_BODY_PART, Name: "FEMUR", ShortName: "FEMUR"}
)

var allColumnValueString = []ColumnValueString{
	// Giới tính
	GENDER_MALE,
	GENDER_FEMALE,
	// Telerad reading order status
	TELERAD_READING_ORDER_STATUS_UNREAD,
	TELERAD_READING_ORDER_STATUS_READING,
	TELERAD_READING_ORDER_STATUS_PENDING_APPROVAL,
	TELERAD_READING_ORDER_STATUS_APPROVED,
	// PACS body part
	PACS_BODY_PART_BRAIN,
	PACS_BODY_PART_HEAD,
	PACS_BODY_PART_LSPINE,
	PACS_BODY_PART_GENERALABDOMEN,
	PACS_BODY_PART_FEMALEPELVIS,
	PACS_BODY_PART_FOREARM,
	PACS_BODY_PART_FOOT,
	PACS_BODY_PART_FA,
	PACS_BODY_PART_ST,
	PACS_BODY_PART_ORBITS,
	PACS_BODY_PART_NECK,
	PACS_BODY_PART_CHEST,
	PACS_BODY_PART_LIVER,
	PACS_BODY_PART_SHOULDER,
	PACS_BODY_PART_WRIST,
	PACS_BODY_PART_ANKLE,
	PACS_BODY_PART_SKULL,
	PACS_BODY_PART_CSPINE,
	PACS_BODY_PART_BREAST,
	PACS_BODY_PART_GENERALPELVIS,
	PACS_BODY_PART_KNEE,
	PACS_BODY_PART_LEG,
	PACS_BODY_PART_CB,
	PACS_BODY_PART_TB,
	PACS_BODY_PART_PITUITARY,
	PACS_BODY_PART_TSPINE,
	PACS_BODY_PART_ABDOMEN,
	PACS_BODY_PART_MALEPELVIS,
	PACS_BODY_PART_HAND,
	PACS_BODY_PART_FEMUR,
}

// NameByValueAndCode trả Name (nhãn hiển thị) của 1 giá trị cột theo (value, code);
// rỗng nếu không khớp.
func NameByValueAndCode(value string, code string) string {
	if cv := FromValueAndCodeString(value, code); cv != nil {
		return cv.Name
	}
	return ""
}

func FromValueAndCodeString(value string, code string) *ColumnValueString {
	for _, v := range allColumnValueString {
		if v.Value == value && v.Code == code {
			return &v
		}
	}

	return nil
}

func GetAllStringTypeByCode(code string) []ColumnValueString {
	output := []ColumnValueString{}

	for _, v := range allColumnValueString {
		if v.Code == code {
			output = append(output, v)
		}
	}

	return output
}
