package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type pos struct {
	x, y int
}

type qItem struct {
	p     pos
	steps int
}

func (i qItem) Less(a qItem) bool {
	return i.steps < a.steps
}

func solvePartOne(input string) (int, error) {
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

	end := pos{x: w - 1, y: h - 1}
	pq := &generic.PriorityQueue[qItem]{}
	heap.Init(pq)
	heap.Push(pq, qItem{p: pos{0, 0}, steps: 0})

	visited := map[pos]bool{}
	for pq.Len() != 0 {
		el := heap.Pop(pq).(qItem)
		if el.p == end {
			return el.steps, nil
		}

		if visited[el.p] {
			continue
		}

		visited[el.p] = true
		for _, dir := range []pos{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1},
		} {
			nextP := pos{x: el.p.x + dir.x, y: el.p.y + dir.y}
			if !bounds(nextP) || grid[nextP.y][nextP.x] == "#" {
				continue
			}

			heap.Push(
				pq,
				qItem{
					p:     nextP,
					steps: el.steps + 1,
				},
			)
		}
	}

	return -1, nil
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Println()
	}
}
