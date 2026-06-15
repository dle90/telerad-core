package utils

import "github.com/google/uuid"

// EqualPtr trả về true nếu hai con trỏ cùng trỏ tới giá trị bằng nhau,
// hoặc cùng là nil. Nếu chỉ một trong hai là nil thì trả về false.
func EqualPtr[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == b
	}

	return *a == *b
}

func GenerateNewUuidV7() (uuid.UUID, error) {
	return uuid.NewV7()
}

// RemoveDuplicates trả về slice mới đã loại bỏ phần tử trùng lặp, giữ nguyên thứ tự
// xuất hiện đầu tiên. Dùng chung cho mọi kiểu comparable (uuid.UUID, string, int, ...).
func RemoveDuplicatesFromArray[T comparable](input []T) []T {
	seen := make(map[T]struct{}, len(input))
	output := make([]T, 0, len(input))

	for _, item := range input {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		output = append(output, item)
	}

	return output
}
