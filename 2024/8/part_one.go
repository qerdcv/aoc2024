package main

import (
	"bufio"
	"io"
)

type pos struct {
	y, x int
}

func solvePartOne(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	var grid [][]byte
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}

	antennas := map[byte][]pos{}
	for y, row := range grid {
		for x, ch := range row {
			if ch == '.' {
				continue
			}

			antennas[ch] = append(antennas[ch], pos{y, x})
		}
	}

	n := len(grid)

	total := 0
	antiNodes := map[pos]struct{}{}
	for _, ant := range antennas {
		for i, a1 := range ant {
			for _, a2 := range ant[i+1:] {
				dy := a2.y - a1.y
				dx := a2.x - a1.x
				p1 := pos{a1.y - dy, a1.x - dx}
				p2 := pos{a2.y + dy, a2.x + dx}

				_, ok := antiNodes[p1]
				if !ok && bounds(p1, n) {
					antiNodes[p1] = struct{}{}
					total += 1
				}

				_, ok = antiNodes[p2]
				if !ok && bounds(p2, n) {
					antiNodes[p2] = struct{}{}
					total += 1
				}
			}
		}
	}

	return total, nil
}

func bounds(p pos, n int) bool {
	return p.x >= 0 && p.x < n && p.y >= 0 && p.y < n
}
