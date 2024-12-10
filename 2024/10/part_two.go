package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(input string) (int, error) {
	grid := generic.Map(strings.Split(input, "\n"), func(ts string) []int {
		return generic.Map(strings.Split(ts, ""), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
	})

	n := len(grid)
	total := 0
	for x := range n {
		for y := range n {
			if grid[x][y] != 0 {
				continue
			}

			total += bfs2(grid, pos{x, y})
		}
	}

	return total, nil
}

func bfs2(grid [][]int, p pos) int {
	q := list.New()
	q.PushBack(p)

	res := 0
	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		p = el.Value.(pos)

		if grid[p.x][p.y] == 9 {
			res++
		}

		for _, ds := range []pos{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		} {
			nextP := pos{x: p.x + ds.x, y: p.y + ds.y}
			if bounds(nextP, len(grid)) && grid[nextP.x][nextP.y] == grid[p.x][p.y]+1 {
				q.PushBack(nextP)
			}
		}
	}

	return res
}
