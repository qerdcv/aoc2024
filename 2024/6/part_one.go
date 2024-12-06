package main

import (
	"fmt"
	"slices"
)

func solvePartOne(lines [][]byte) (int, error) {
	y, x, dir := getGuardPos(lines)
	visited := map[string]struct{}{fmt.Sprintf("%dx%d", y, x): {}}

	for {
		dy, dx := deltaFromDir(dir)
		y += dy
		x += dx
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
			break
		}

		if lines[y][x] == '#' {
			y, x = y-dy, x-dx
			dir = (dir + 1) % 4
			continue
		}

		visited[fmt.Sprintf("%dx%d", y, x)] = struct{}{}
	}

	return len(visited), nil
}

func deltaFromDir(dir int) (int, int) {
	switch dir {
	case up:
		return -1, 0
	case right:
		return 0, 1
	case down:
		return 1, 0
	case left:
		return 0, -1
	}

	return 0, 0
}

func getGuardPos(lines [][]byte) (int, int, int) {
	for y, row := range lines {
		for x, ch := range row {
			if !slices.Contains([]byte{'>', '<', '^', 'v'}, ch) {
				continue
			}

			switch ch {
			case '>':
				return y, x, right
			case '<':
				return y, x, left
			case '^':
				return y, x, up
			case 'v':
				return y, x, down
			}
		}
	}

	return 0, 0, -1
}
