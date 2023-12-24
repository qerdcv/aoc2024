package generic

type List[T any] []T

func (l *List[T]) PopStart() T {
	if len(*l) == 0 {
		return *new(T)
	}

	item := (*l)[0]
	*l = (*l)[1:]
	return item
}

func (l *List[T]) Pop() T {
	if len(*l) == 0 {
		return *new(T)
	}

	lastIdx := len(*l) - 1
	item := (*l)[lastIdx]
	*l = (*l)[:lastIdx]
	return item
}

func (l *List[T]) Append(val T) {
	*l = append(*l, val)
}
