package day11

import (
	"bufio"
	"io"
)

type coord struct {
	x, y int
}

func ResolvePartOne(r io.Reader) int {
	universe := parseTheUniverse(r)
	coords := findGalaxyCoords(universe)
	rows, cols := getExpansionCorrelations(universe)
	return calculateShortestPaths(expandCoords(coords, 2, rows, cols))
}

func ResolvePartTwo(r io.Reader) int {
	universe := parseTheUniverse(r)
	coords := findGalaxyCoords(universe)
	rows, cols := getExpansionCorrelations(universe)
	return calculateShortestPaths(expandCoords(coords, 1_000_000, rows, cols))
}

func calculateShortestPaths(coords []coord) int {
	pathsLen := 0

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]

			pathsLen += abs(c1.x-c2.x) + abs(c1.y-c2.y)
		}
	}

	return pathsLen
}

func expandCoords(coords []coord, mul int, rows, cols []int) []coord {
	mul -= 1

	dr := 0
	for _, row := range rows {
		row += dr
		for idx := range coords {
			if coords[idx].y > row {
				coords[idx].y += mul
			}
		}

		dr += mul
	}

	dc := 0
	for _, col := range cols {
		col += dc
		for idx := range coords {
			if coords[idx].x > col {
				coords[idx].x += mul
			}
		}

		dc += mul
	}

	return coords
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func findGalaxyCoords(u [][]byte) []coord {
	var coords []coord

	for y, row := range u {
		for x, col := range row {
			if col == '#' {
				coords = append(coords, coord{x, y})
			}
		}
	}

	return coords
}

func getExpansionCorrelations(u [][]byte) ([]int, []int) {
	var (
		rowsWithoutGalaxies []int
		colsWithoutGalaxies []int
	)

rowLoop:
	for rowIdx, row := range u {
		for _, ch := range row {
			if ch == '#' {
				continue rowLoop
			}
		}
		rowsWithoutGalaxies = append(rowsWithoutGalaxies, rowIdx)
	}

	colLen := len(u[0])

colLoop:
	for i := 0; i < colLen; i++ {
		for _, row := range u {
			if row[i] == '#' {
				continue colLoop
			}
		}

		colsWithoutGalaxies = append(colsWithoutGalaxies, i)
	}
	return rowsWithoutGalaxies, colsWithoutGalaxies
}

func parseTheUniverse(r io.Reader) [][]byte {
	s := bufio.NewScanner(r)
	var res [][]byte
	for s.Scan() {
		res = append(res, []byte(s.Text()))
	}

	return res
}
