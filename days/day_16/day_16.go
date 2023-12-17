package day16

import (
	"bufio"
	"io"
	"math"
)

var visited = map[vec]bool{}

type point struct {
	x, y int
}
type vec struct {
	point      // displays current position
	dx, dy int // displays current vector of movement
}

func ResolvePartOne(r io.Reader) int {
	f := parseField(r)

	energizeField(f, vec{
		point: point{x: 0, y: 0},
		dx:    1,
		dy:    0,
	}) // start from left top corner and move to the right

	// need to make unique occupied cells
	energizedCells := map[point]bool{}

	for v := range visited {
		energizedCells[v.point] = true
	}

	return len(energizedCells)
}

func ResolvePartTwo(r io.Reader) int {
	f := parseField(r)

	rCnt := len(f)
	cCnt := len(f[0])

	maxEnergizedCells := math.MinInt64
	// try all left to right combinations
	for i := 0; i < rCnt; i++ {
		clear(visited)

		energizeField(f, vec{
			point: point{x: 0, y: i},
			dx:    1,
			dy:    0,
		}) // start from left top corner and move to the right

		// need to make unique occupied cells
		energizedCells := map[point]bool{}

		for v := range visited {
			energizedCells[v.point] = true
		}

		maxEnergizedCells = max(maxEnergizedCells, len(energizedCells))
	}

	// try all top to bottom combinations
	for i := 0; i < cCnt; i++ {
		clear(visited)

		energizeField(f, vec{
			point: point{x: i, y: 0},
			dx:    0,
			dy:    1,
		}) // start from left top corner and move to the right

		// need to make unique occupied cells
		energizedCells := map[point]bool{}

		for v := range visited {
			energizedCells[v.point] = true
		}

		maxEnergizedCells = max(maxEnergizedCells, len(energizedCells))
	}

	// try all right to left combinations
	for i := rCnt - 1; i >= 0; i-- {
		clear(visited)

		energizeField(f, vec{
			point: point{x: cCnt - 1, y: i},
			dx:    -1,
			dy:    0,
		}) // start from left top corner and move to the right

		// need to make unique occupied cells
		energizedCells := map[point]bool{}

		for v := range visited {
			energizedCells[v.point] = true
		}

		maxEnergizedCells = max(maxEnergizedCells, len(energizedCells))
	}

	// try all bottom to top combinations
	for i := cCnt - 1; i >= 0; i-- {
		clear(visited)

		energizeField(f, vec{
			point: point{x: i, y: rCnt},
			dx:    0,
			dy:    -1,
		}) // start from left top corner and move to the right

		// need to make unique occupied cells
		energizedCells := map[point]bool{}

		for v := range visited {
			energizedCells[v.point] = true
		}

		maxEnergizedCells = max(maxEnergizedCells, len(energizedCells))
	}

	return maxEnergizedCells
}

func energizeField(f [][]byte, v vec) {
	for {
		if !((0 <= v.x && v.x < len(f[0])) && (0 <= v.y && v.y < len(f))) {
			return
		}

		if visited[v] {
			return
		}

		visited[v] = true
		switch f[v.y][v.x] {
		case '.':
		case '|':
			if v.dy == 0 {
				energizeField(f, vec{
					point: v.point,
					dy:    -1,
					dx:    0,
				})
				energizeField(f, vec{
					point: v.point,
					dy:    1,
					dx:    0,
				})
				return
			}
		case '-':
			if v.dx == 0 {
				// if light come from top or bottom - it will be split by X axe
				energizeField(f, vec{
					point: v.point,
					dy:    0,
					dx:    1,
				})
				energizeField(f, vec{
					point: v.point,
					dy:    0,
					dx:    -1,
				})
				return
			}

		case '/':
			if v.dx != 0 {
				switch v.dx {
				case 1: // from right to top
					v.dx = 0
					v.dy = -1
				case -1: // from left to bottom
					v.dx = 0
					v.dy = 1
				default:
					panic("unreachable")
				}
			} else {
				// /
				switch v.dy {
				case 1: // from top to left
					v.dy = 0
					v.dx = -1
				case -1: // from bottom to right
					v.dy = 0
					v.dx = 1
				default:
					panic("unreachable")
				}
			}
		case '\\':
			if v.dx != 0 {
				switch v.dx {
				case 1: // from right to bottom
					v.dx = 0
					v.dy = 1
				case -1: // from left to top
					v.dx = 0
					v.dy = -1
				default:
					panic("unreachable")
				}
			} else {
				// \
				switch v.dy {
				case 1: // from top to right
					v.dy = 0
					v.dx = 1
				case -1: // from bottom to left
					v.dy = 0
					v.dx = -1
				default:
					panic("unreachable")
				}
			}
		default:
			panic("unreachable")
		}

		v.y += v.dy
		v.x += v.dx
	}
}

func parseField(r io.Reader) [][]byte {
	s := bufio.NewScanner(r)
	var field [][]byte
	for s.Scan() {
		field = append(field, []byte(s.Text()))
	}

	return field
}
