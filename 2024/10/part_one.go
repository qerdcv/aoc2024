package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type pos struct {
	x, y int
}

func solvePartOne(input string) (int, error) {
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

			total += bfs(grid, pos{x, y})
		}
	}

	return total, nil
}

func bfs(grid [][]int, p pos) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid))
	}

	q := list.New()
	q.PushBack(p)
	visited[p.x][p.y] = true

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
			if bounds(nextP, len(grid)) && !visited[nextP.x][nextP.y] && grid[nextP.x][nextP.y] == grid[p.x][p.y]+1 {
				visited[nextP.x][nextP.y] = true
				q.PushBack(nextP)
			}
		}
	}

	return res
}

func bounds(p pos, n int) bool {
	return p.x >= 0 && p.x < n && p.y >= 0 && p.y < n
}
