package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type pos struct {
	x, y int
}

func solvePartOne(input string) (int, error) {
	l := generic.Map(strings.Split(input, "\n"), func(t string) [2]pos {
		rp, dp := parseInp(t)
		return [2]pos{rp, dp}
	})

	w := 101
	h := 103

	for range 100 {
		for i := range l {
			l[i][0].x = l[i][0].x + l[i][1].x
			l[i][0].y = l[i][0].y + l[i][1].y

			if l[i][0].x < 0 {
				l[i][0].x = w + l[i][0].x
			}

			if l[i][0].y < 0 {
				l[i][0].y = h + l[i][0].y
			}

			l[i][0].x %= w
			l[i][0].y %= h
		}
	}

	hw := w / 2
	hh := h / 2

	qs := [4]int{0, 0, 0, 0}
	for _, r := range l {
		p := r[0]

		if p.x == hw || p.y == hh {
			continue
		}

		q := 0
		if p.x >= hw {
			q = 1
		}

		if p.y > hh {
			q += 2
		}

		qs[q]++
	}

	return qs[0] * qs[1] * qs[2] * qs[3], nil
}

func printTitles(wr io.Writer, w, h int, l [][2]pos) {
	for y := range h {
		for x := range w {
			robotCnt := 0
			for _, r := range l {
				if r[0].x == x && r[0].y == y {
					robotCnt++
				}
			}

			if robotCnt == 0 {
				fmt.Fprint(wr, ".")
			} else {
				fmt.Fprint(wr, robotCnt)
			}
		}
		fmt.Fprint(wr, "\n")
	}

	fmt.Fprintln(wr)
}

func parseInp(line string) (pos, pos) {
	parts := strings.Split(line, " ")
	rPosParts := strings.Split(parts[0], ",")
	rP := pos{}
	rP.x, _ = strconv.Atoi(rPosParts[0][2:])
	rP.y, _ = strconv.Atoi(rPosParts[1])

	dPParts := strings.Split(parts[1], ",")
	dP := pos{}
	dP.x, _ = strconv.Atoi(dPParts[0][2:])
	dP.y, _ = strconv.Atoi(dPParts[1])

	return rP, dP
}
