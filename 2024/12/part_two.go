package main

import (
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type edge struct {
	p  pos
	dp pos
}

var dps = []pos{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func bounds(p pos, gridSize int) bool {
	return p.y >= 0 && p.y < gridSize && p.x >= 0 && p.x < gridSize
}

func scanRegion(p pos, ch string, grid [][]string, visited [][]bool, edges map[edge]struct{}, gridSize int) int {
	if visited[p.y][p.x] || grid[p.y][p.x] != ch {
		return 0
	}

	visited[p.y][p.x] = true
	area := 1

	for _, dp := range dps {
		newP := pos{y: p.y + dp.y, x: p.x + dp.x}

		if !bounds(newP, gridSize) || grid[newP.y][newP.x] != ch {
			edges[edge{p: p, dp: dp}] = struct{}{}
			continue
		}
		area += scanRegion(newP, ch, grid, visited, edges, gridSize)
	}

	return area
}

func solvePartTwo(input string) (int, error) {
	grid := generic.Map(strings.Split(input, "\n"), func(t string) []string {
		return strings.Split(t, "")
	})
	gridSize := len(grid)

	visited := make([][]bool, gridSize)
	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}

	total := 0
	for y, row := range grid {
		for x, ch := range row {
			p := pos{y: y, x: x}
			if visited[p.y][p.x] {
				continue
			}

			edges := make(map[edge]struct{})
			area := scanRegion(p, string(ch), grid, visited, edges, gridSize)

			dupEdges := 0
			for e := range edges {
				reverseEdge := edge{p: pos{y: e.p.y + e.dp.x, x: e.p.x + e.dp.y}, dp: e.dp}
				if _, exists := edges[reverseEdge]; exists {
					dupEdges++
				}
			}

			total += area * (len(edges) - dupEdges)
		}
	}

	return total, nil
}
