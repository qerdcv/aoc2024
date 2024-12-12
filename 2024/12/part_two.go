package main

import (
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(input string) (int, error) {
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

	var scanRegion func(p pos, ch string, edges map[pos]struct{}) int

	scanRegion = func(p pos, ch string, edges map[pos]struct{}) int {
		if visited[p.y][p.x] || grid[p.y][p.x] != ch {
			return 0
		}

		visited[p.y][p.x] = true
		area := 1
		for _, dp := range []pos{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		} {
			newP := pos{y: p.y + dp.y, x: p.x + dp.x}
			if !bound(newP) || grid[newP.y][newP.x] != ch {
				edges[newP] = struct{}{}
				continue
			}

			area += scanRegion(newP, ch, edges)
		}

		return area
	}

	total := 0
	for y, row := range grid {
		for x, ch := range row {
			p := pos{y: y, x: x}
			if visited[p.y][p.x] {
				continue
			}

			// TODO: count sides
			edges := map[pos]struct{}{}
			a := scanRegion(p, ch, edges)
			total += a
		}
	}

	return total, nil
}
