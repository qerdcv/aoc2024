package main

import (
	"fmt"
	"io"
)

func solvePartTwo(r io.Reader) (int, error) {
	lines := readLines(r)
	y, x, dy, dx := getGuardPos(lines)

	total := 0
	for localY := range len(lines) {
		for localX := range len(lines) {
			if ch := lines[localY][localX]; ch == '#' || (localY == y && localX == x) {
				continue
			}

			tmp := lines[localY][localX]
			lines[localY][localX] = '#'

			if isLooped(lines, y, x, dy, dx) {
				total += 1
			}

			lines[localY][localX] = tmp
		}
	}

	for _, row := range lines {
		fmt.Println(string(row))
	}

	return total, nil
}

func isLooped(lines [][]byte, y, x, dy, dx int) bool {
	visited := map[string]struct{}{}

	for {
		y, x = y+dy, x+dx
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
			return false
		}

		if _, ok := visited[getVisitedKey(y, x, dy, dx)]; ok {
			return true
		}

		if lines[y][x] != '#' && lines[y][x] != '^' {
			if dx != 0 {
				lines[y][x] = '-'
			}

			if dy != 0 {
				lines[y][x] = '|'
			}
		}

		if lines[y][x] == '#' {
			y, x = y-dy, x-dx
			lines[y][x] = '+'
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

		visited[getVisitedKey(y, x, dy, dx)] = struct{}{}
	}
}

func getVisitedKey(y, x, dy, dx int) string {
	return fmt.Sprintf("%dx%dx%dx%d", y, x, dy, dx)
}
