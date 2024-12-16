package main

import (
	"container/heap"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type dir int

const (
	up dir = iota
	right
	down
	left
)

func (d dir) toPos() pos {
	switch d {
	case up:
		return pos{x: 0, y: -1}
	case right:
		return pos{x: 1, y: 0}
	case down:
		return pos{x: 0, y: 1}
	case left:
		return pos{x: -1, y: 0}
	}

	panic("unknown direction")
}

type pos struct {
	x, y int
}

func (p pos) move(d dir) pos {
	dp := d.toPos()
	return pos{x: p.x + dp.x, y: p.y + dp.y}
}

type qItem struct {
	p     pos
	score int
	d     dir
}

type pq []qItem

// Len implements heap.Interface.
func (p pq) Len() int {
	return len(p)
}

// Less implements heap.Interface.
func (p pq) Less(i int, j int) bool {
	return p[i].score < p[j].score
}

// Swap implements heap.Interface.
func (p pq) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// Pop implements heap.Interface.
func (p *pq) Pop() any {
	pLen := p.Len() - 1
	val := (*p)[pLen]
	*p = (*p)[:pLen]
	return val
}

// Push implements heap.Interface.
func (p *pq) Push(x any) {
	(*p) = append((*p), x.(qItem))
}

type cacheKey struct {
	p pos
	d dir
}

func solvePartOne(input string) (int, error) {
	grid := generic.Map(strings.Split(input, "\n"), func(t string) []string {
		return strings.Split(t, "")
	})

	start := findStart(grid, "S")
	q := &pq{}
	heap.Init(q)

	heap.Push(q, qItem{
		p:     start,
		score: 1000,
		d:     up,
	})

	cache := map[cacheKey]int{}

	for q.Len() != 0 {
		el := heap.Pop(q).(qItem)
		currentP := el.p
		currentD := el.d
		currentS := el.score

		if grid[currentP.y][currentP.x] == "E" {
			return currentS, nil
		}

		ck := cacheKey{
			p: currentP,
			d: currentD,
		}

		if existingScore, exists := cache[ck]; exists && existingScore <= currentS {
			continue
		}

		cache[ck] = currentS

		for _, nextD := range []dir{
			currentD,
			(currentD + 1) % 4,
			(currentD + 3) % 4,
		} {
			nextP := currentP.move(nextD)

			if grid[nextP.y][nextP.x] == "#" {
				continue
			}

			nextS := currentS + 1
			if nextD != currentD {
				nextS += 1000
			}

			if cachedS, found := cache[cacheKey{p: nextP, d: nextD}]; !found || cachedS > nextS {
				heap.Push(q, qItem{
					nextP,
					nextS,
					nextD,
				})
			}
		}
	}

	return 0, nil
}

func findStart(grid [][]string, el string) pos {
	for y, row := range grid {
		for x, col := range row {
			if col == el {
				return pos{x, y}
			}
		}
	}

	panic("can't find start")
}
