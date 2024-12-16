package main

import (
	"container/heap"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type cackeVal struct {
	score   int
	prevPos pos
	prevDir dir
}

// 10039 - too high
func solvePartTwo(input string) (int, error) {
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

	cache := map[cacheKey]cackeVal{}

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
