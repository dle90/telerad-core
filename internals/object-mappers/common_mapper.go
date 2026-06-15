package objectMappers

// emptyIfNil trả về slice rỗng thay vì nil để JSON serialize ra [] thay vì null.
func emptyIfNil[T any](s []T) []T {
	if s == nil {
		return []T{}
	}
	return s
}
