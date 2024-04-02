package gozod

type Optional[T any] struct{ value *T }

func NewEmptyOptional[T any]() Optional[T] {
	return Optional[T]{}
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o Optional[T]) UnWrap() T {
	return *o.value
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.IsPresent() {
		return o.UnWrap()
	}
	return defaultValue
}
