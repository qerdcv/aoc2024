package day17

import (
	"bufio"
	"io"

	"github.com/qerdcv/aoc2023/internal/generic"
)

type coord struct {
	x, y int
}

type node struct {
	coord
	val int
}

type dir struct {
	dx, dy int
}

func ResolvePartOne(r io.Reader) int {
	cm := parseCM(r)

	return 0
}

func ResolvePartTwo(r io.Reader) int {
	return 0
}

func findMinHeatLoss(cm [][]int) int {
	visited := map[coord]node{}
	columns := len(cm[0])
	rows := len(cm)

	queue := generic.List[node]{
		{coord: coord{0, 0}, val: cm[rows-1][columns-1]},
	}
	dirs := []dir{
		{0, 1}, {0, -1},
		{1, 0}, {-1, 0},
	}

	for len(queue) != 0 {
		n := queue.PopStart()
		for _, d := range dirs {
			dx := n.x + d.dx
			dy := n.y + d.dy

			if 0 < dx 
		}
	}
	return 0
}

func parseCM(r io.Reader) [][]int {
	s := bufio.NewScanner(r)
	var res [][]int

	for s.Scan() {
		rawLine := s.Text()
		line := make([]int, len(rawLine))
		for idx, ch := range rawLine {
			line[idx] = int(ch - '0')
		}

		res = append(res, line)
	}

	return res
}
