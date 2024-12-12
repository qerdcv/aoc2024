package main

import (
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type pos struct {
	x, y int
}

func solvePartOne(input string) (int, error) {
	grid := generic.Map(strings.Split(input, "\n"), func(t string) []string {
		return strings.Split(t, "")
	})
	gL := len(grid)

	visited := make([][]bool, gL)
	for i := range visited {
		visited[i] = make([]bool, gL)
	}

	bound := func(p pos) bool {
		return p.y < gL && p.y >= 0 && p.x < gL && p.x >= 0
	}

	neighborCount := func(p pos, ch string) int {
		n := 0
		for _, dp := range []pos{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		} {
			newP := pos{y: p.y + dp.y, x: p.x + dp.x}
			if !bound(newP) {
				continue
			}

			if ch == grid[newP.y][newP.x] {
				n++
			}
		}

		return n
	}

	var scanRegion func(p pos, ch string) (int, int)

	scanRegion = func(p pos, ch string) (int, int) {
		if visited[p.y][p.x] || grid[p.y][p.x] != ch {
			return 0, 0
		}

		visited[p.y][p.x] = true
		res := 4 - neighborCount(p, ch)
		area := 1
		for _, dp := range []pos{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		} {
			newP := pos{y: p.y + dp.y, x: p.x + dp.x}
			if !bound(newP) || grid[newP.y][newP.x] != ch {
				continue
			}
			r2, a2 := scanRegion(newP, ch)
			res += r2
			area += a2
		}

		return res, area
	}

	total := 0
	for y, row := range grid {
		for x, ch := range row {
			p := pos{y: y, x: x}
			if visited[p.y][p.x] {
				continue
			}

			a, b := scanRegion(p, ch)
			total += a * b
		}
	}

	return total, nil
}
