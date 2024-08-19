package pointer

func Empty[T any]() T {
	var zero T
	return zero
}
func FromPtr[T any](ptr *T) T {
	if ptr == nil {
		return Empty[T]()
	}

	return *ptr
}

func ToPtr[T any](v T) *T {
	return &v
}
