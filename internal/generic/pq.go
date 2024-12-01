package generic

import "container/heap"

type PriorityQueue []int

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *PriorityQueue) Pop() any {
	pLen := p.Len() - 1
	val := (*p)[pLen]
	*p = (*p)[:pLen]
	return val
}

var _ heap.Interface = (*PriorityQueue)(nil)
