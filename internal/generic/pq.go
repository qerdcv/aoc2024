package generic

type comparable[T any] interface {
	Less(T) bool
}

type PriorityQueue[T comparable[T]] struct {
	items []T
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.items)
}

func (p *PriorityQueue[T]) Less(i, j int) bool {
	return p.items[i].Less(p.items[j])
}

func (p *PriorityQueue[T]) Swap(i, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
}

func (p *PriorityQueue[T]) Push(x any) {
	p.items = append(p.items, x.(T))
}

func (p *PriorityQueue[T]) Pop() any {
	pLen := p.Len() - 1
	val := p.items[pLen]
	p.items = p.items[:pLen]
	return val
}
