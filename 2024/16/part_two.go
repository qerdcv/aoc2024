package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type qItem2 struct {
	p     pos
	score int
	steps int
	d     dir
	prev  *qItem2
}

func (i qItem2) Less(a qItem2) bool {
	return i.score < a.score
}

func solvePartTwo(input string) (int, error) {
	grid := generic.Map(strings.Split(input, "\n"), func(t string) []string {
		return strings.Split(t, "")
	})

	start := findEl(grid, "S")
	end := findEl(grid, "E")
	q := &generic.PriorityQueue[qItem2]{}
	heap.Init(q)

	heap.Push(q, qItem2{
		p:     start,
		score: 1000,
		steps: 0,
		d:     up,
		prev:  nil,
	})

	cache := map[cacheKey]int{}
	best := -1
	var items []qItem2
	for q.Len() != 0 {
		el := heap.Pop(q).(qItem2)
		currentP := el.p
		currentD := el.d
		currentS := el.score
		currentSteps := el.steps

		if currentP == end {
			if best == -1 {
				best = currentS
			}

			if currentS == best {
				items = append(items, el)
			}
		}

		ck := cacheKey{
			p: currentP,
			d: currentD,
		}

		if v, ok := cache[ck]; ok && v < currentS {
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

			heap.Push(q, qItem2{
				nextP,
				nextS,
				currentSteps + 1,
				nextD,
				&el,
			})
		}
	}

	coords := map[pos]struct{}{}
	for _, el := range items {
		coords[el.p] = struct{}{}
		current := el.prev
		for current != nil {
			coords[current.p] = struct{}{}
			current = current.prev
		}
	}

	return len(coords), nil
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
	fmt.Println()
}
