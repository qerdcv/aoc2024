package main

import (
	"bufio"
	"io"
)

func solvePartOne(r io.Reader) (int, error) {
	lines := readLines(r)
	total := 0
	xCnt := 0
	for y, line := range lines {
		for x, ch := range line {
			if ch != 'X' {
				continue
			}
			xCnt++
			for _, b := range []bool{
				scanXMAS(y, x, 1, 0, lines),
				scanXMAS(y, x, 0, 1, lines),
				scanXMAS(y, x, -1, 0, lines),
				scanXMAS(y, x, 0, -1, lines),
				scanXMAS(y, x, 1, 1, lines),
				scanXMAS(y, x, 1, -1, lines),
				scanXMAS(y, x, -1, 1, lines),
				scanXMAS(y, x, -1, -1, lines),
			} {
				if b {
					total++
				}
			}
		}
	}

	return total, nil
}

func scanXMAS(y, x, dy, dx int, lines [][]byte) bool {
	yL := len(lines)
	xL := len(lines[0])

	word := []byte{'X', 'M', 'A', 'S'}
	matchCnt := 0
	for range 4 {
		if lines[y][x] != word[matchCnt] {
			return false
		}

		y += dy
		x += dx
		matchCnt++

		if x >= xL || y >= yL || x < 0 || y < 0 {
			break
		}
	}

	if matchCnt != 4 {
		return false
	}

	return true
}
func readLines(r io.Reader) [][]byte {
	s := bufio.NewScanner(r)
	var lines [][]byte
	for s.Scan() {
		lines = append(lines, []byte(s.Text()))
	}

	return lines
}
