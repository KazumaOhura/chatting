package nullable

type Nullable[T any] struct {
	Value   T
	IsValid bool
}
