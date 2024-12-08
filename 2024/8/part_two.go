package main

import (
	"bufio"
	"io"
)

func solvePartTwo(r io.Reader) (int, error) {
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
				dx := a2.x - a1.x
				dy := a2.y - a1.y
				p1, p2 := pos{a1.y, a1.x}, pos{a2.y, a2.x}
				for bounds(p1, n) || bounds(p2, n) {
					if bounds(p1, n) {
						if _, ok := antiNodes[p1]; !ok {
							antiNodes[p1] = struct{}{}
							total += 1
						}
					}

					if bounds(p2, n) {
						if _, ok := antiNodes[p2]; !ok {
							antiNodes[p2] = struct{}{}
							total += 1
						}
					}
					p1, p2 = pos{p1.y - dy, p1.x - dx}, pos{p2.y + dy, p2.x + dx}
				}
			}
		}
	}

	return total, nil
}
