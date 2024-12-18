package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type qItem2 struct {
	p     pos
	steps int
}

func (i qItem2) Less(a qItem2) bool {
	return i.steps < a.steps
}

func solvePartTwo(input string) (int, error) {
	var w, h int
	var coords []pos
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		coord := pos{}
		coord.x, _ = strconv.Atoi(parts[0])
		coord.y, _ = strconv.Atoi(parts[1])
		coords = append(coords, coord)

		if coord.x >= w {
			w = coord.x + 1
		}

		if coord.y >= h {
			h = coord.y + 1
		}
	}

	grid := make([][]string, h)
	for i := range grid {
		grid[i] = make([]string, w)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, coord := range coords[:1024] {
		grid[coord.y][coord.x] = "#"
	}

	printGrid(grid)

	bounds := func(p pos) bool {
		return p.x >= 0 && p.x < w && p.y >= 0 && p.y < h
	}

	start := pos{x: 0, y: 0}
	end := pos{x: w - 1, y: h - 1}

	findPath := func() bool {
		visited := map[pos]bool{}
		q := list.New()
		q.PushBack(start)
		for q.Len() != 0 {
			el := q.Front()
			q.Remove(el)
			p := el.Value.(pos)

			if p == end {
				return true
			}

			if visited[p] {
				continue
			}

			visited[p] = true
			for _, dir := range []pos{
				{-1, 0},
				{0, 1},
				{1, 0},
				{0, -1},
			} {
				nextP := pos{x: p.x + dir.x, y: p.y + dir.y}
				if !bounds(nextP) || grid[nextP.y][nextP.x] == "#" {
					continue
				}

				q.PushBack(nextP)
			}
		}

		return false
	}

	for _, coord := range coords[1024:] {
		grid[coord.y][coord.x] = "#"
		if !findPath() {
			fmt.Println(coord)
			break
		}
	}

	return -1, nil
}
