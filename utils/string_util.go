package utils

import (
	"bytes"
	"crypto/rand"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	commonUtils "github.com/BeeTechHub/go-common/utils"
	"golang.org/x/text/unicode/norm"
)

// Mang cac ky tu goc co dau
var SOURCE_CHARACTERS, LL_LENGTH = stringToRune(`ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự`)

// Mang cac ky tu thay the khong dau
var DESTINATION_CHARACTERS, _ = stringToRune(`AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUu`)

func stringToRune(s string) ([]string, int) {
	ll := utf8.RuneCountInString(s)
	var texts = make([]string, ll+1)
	var index = 0
	for _, runeValue := range s {
		texts[index] = string(runeValue)
		index++
	}
	return texts, ll
}

func binarySearch(sortedArray []string, key string, low int, high int) int {
	var middle int = (low + high) / 2
	if high < low {
		return -1
	}
	if key == sortedArray[middle] {
		return middle
	} else if key < sortedArray[middle] {
		return binarySearch(sortedArray, key, low, middle-1)
	} else {
		return binarySearch(sortedArray, key, middle+1, high)
	}
}

/** * Bo dau 1 ky tu * * @param ch * @return */
func removeAccentChar(ch string) string {
	var index int = binarySearch(SOURCE_CHARACTERS, ch, 0, LL_LENGTH)
	if index >= 0 {
		ch = DESTINATION_CHARACTERS[index]
	}
	return ch
}

/** * Bo dau 1 chuoi * * @param s * @return */
func RemoveAccent(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		buffer.WriteString(removeAccentChar(string(runeValue)))
	}
	return buffer.String()
}

/** * Bo dau 1 chuoi * * @param s * @return */
func GenerateCode(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		buffer.WriteString(removeAccentChar(string(runeValue)))
	}
	return strings.ReplaceAll(buffer.String(), " ", "-") + "-" + strconv.FormatInt(time.Now().Unix(), 10)
}

func GenOtp(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func IsValidPassword(s string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	commonUtils.UnUsed(hasSpecial)
	return hasUpper && hasLower && hasNumber
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	isValid := regex.MatchString(email)
	return isValid
}

// Kiểm tra chuỗi có phải là số điện thoại hợp lệ (10 chữ số, bắt đầu bằng 0)
func IsValidPhoneNumber(phone string) bool {
	// Bắt đầu bằng 0 và theo sau là 9 chữ số
	pattern := `^0[0-9]{9}$`
	regex := regexp.MustCompile(pattern)
	isValid := regex.MatchString(phone)
	return isValid
}

// https://zetcode.com/golang/string-format/
// func FormatString(format string, a ...any) string {
// 	return fmt.Sprintf(format, a)
// }

// -1 if v1 is smaller than v2.
// 1 if v1 is greater than v2.
// 0 if both versions are equal.
func CompareVersions(v1, v2 string) int {
	// Split the versions by dot separator
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	// Determine the maximum length
	maxLen := len(v1Parts)
	if len(v2Parts) > maxLen {
		maxLen = len(v2Parts)
	}

	// Compare each corresponding part
	for i := 0; i < maxLen; i++ {
		var v1Num, v2Num int

		// If one version has fewer parts, assume missing parts are zero
		if i < len(v1Parts) {
			v1Num, _ = strconv.Atoi(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2Num, _ = strconv.Atoi(v2Parts[i])
		}

		if v1Num < v2Num {
			return -1
		} else if v1Num > v2Num {
			return 1
		}
	}

	// Versions are equal
	return 0
}

func BuildCode(prefix string, number int64, numberPad int) string {
	numberStr := strconv.FormatInt(number, 10)
	if numberPad > len(numberStr) {
		numberStr = strings.Repeat("0", numberPad-len(numberStr)) + numberStr
	}
	return prefix + numberStr
}

// ToCanonical chuẩn hoá chuỗi để index / so sánh không phân biệt dấu tiếng Việt:
//  1. NFD decompose (tách base char + dấu thanh/dấu mũ/dấu móc).
//  2. Strip combining marks (Unicode category Mn) → bỏ toàn bộ dấu.
//  3. Map đ/Đ → d/D (đặc biệt, không decompose qua NFD).
//  4. UPPERCASE.
//  5. Trim + collapse multi-whitespace về 1 space.
//
// Ví dụ: "  Nguyễn   Văn  Đức  " → "NGUYEN VAN DUC".
func ToCanonical(input string) string {
	decomposed := norm.NFD.String(input)

	var b strings.Builder
	b.Grow(len(decomposed))
	for _, r := range decomposed {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		switch r {
		case 'đ':
			b.WriteRune('d')
		case 'Đ':
			b.WriteRune('D')
		default:
			b.WriteRune(r)
		}
	}

	return strings.ToUpper(strings.Join(strings.Fields(b.String()), " "))
}
