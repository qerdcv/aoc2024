package main

import (
	"io"
)

func solvePartTwo(r io.Reader) (int, error) {
	lines := readLines(r)
	total := 0
	for y, line := range lines {
		for x, ch := range line {
			if ch != 'A' {
				continue
			}

			if scanXmas(y, x, lines) {
				total += 1
			}
		}
	}

	return total, nil
}

func scanXmas(y, x int, lines [][]byte) bool {
	if y < 1 || x < 1 || y == len(lines)-1 || x == len(lines[0])-1 {
		return false
	}

	for _, dc := range [][]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	} {
		dy, dx := dc[0], dc[1]
		if lines[y+dy][x+dx] == 'M' && lines[y-dy][x-dx] == 'S' {
			if (lines[y-dy][x+dx] == 'M' && lines[y+dy][x-dx] == 'S') ||
				(lines[y-dy][x+dx] == 'S' && lines[y+dy][x-dx] == 'M') {
				return true
			}
		}
	}

	return false
}
