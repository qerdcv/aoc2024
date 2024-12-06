package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

func solvePartOne(r io.Reader) (int, error) {
	lines := readLines(r)
	y, x, dy, dx := getGuardPos(lines)
	visited := map[string]struct{}{}

	for {
		y, x = y+dy, x+dx
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
			break
		}

		if lines[y][x] == '#' {
			y, x = y-dy, x-dx
			if dx == 1 {
				dy, dx = 1, 0
			} else if dx == -1 {
				dy, dx = -1, 0
			} else if dy == -1 {
				dy, dx = 0, 1
			} else if dy == 1 {
				dy, dx = 0, -1
			}
			continue
		}

		lines[y][x] = 'X'

		visited[fmt.Sprintf("%dx%d", y, x)] = struct{}{}
	}

	return len(visited), nil
}

func getGuardPos(lines [][]byte) (int, int, int, int) {
	for y, row := range lines {
		for x, ch := range row {
			if !slices.Contains([]byte{'>', '<', '^', 'v'}, ch) {
				continue
			}

			switch ch {
			case '>':
				return y, x, 0, 1
			case '<':
				return y, x, 0, -1
			case '^':
				return y, x, -1, 0
			case 'v':
				return y, x, 1, 0
			}
		}
	}

	return 0, 0, 0, 0
}

func readLines(r io.Reader) [][]byte {
	var lines [][]byte
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, []byte(s.Text()))
	}

	return lines
}
