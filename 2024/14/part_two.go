package main

import (
	"container/list"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(input string) (int, error) {
	l := generic.Map(strings.Split(input, "\n"), func(t string) [2]pos {
		rp, dp := parseInp(t)
		return [2]pos{rp, dp}
	})

	w := 101
	h := 103

	seconds := 1
	for {
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

		if solidity := getMaxRobotsSolidity(l); solidity > 100 {
			break
		}

		seconds++
	}

	return seconds, nil
}

func getMaxRobotsSolidity(l [][2]pos) int {
	positions := make(map[pos]bool, len(l))
	for _, r := range l {
		p := r[0]
		positions[p] = true
	}

	maxSolidity := 0
	seen := make(map[pos]bool, len(positions))
	for p := range positions {
		if seen[p] {
			continue
		}

		solidity := 0
		q := list.New()
		q.PushBack(p)

		for q.Len() != 0 {
			el := q.Front()
			q.Remove(el)
			p := el.Value.(pos)
			if seen[p] {
				continue
			}

			seen[p] = true
			solidity++
			for _, dp := range []pos{
				{-1, 0},
				{0, -1},
				{1, 0},
				{0, 1},
			} {
				newP := pos{x: p.x + dp.x, y: p.y + dp.y}
				if !positions[newP] || seen[newP] {
					continue
				}
				q.PushBack(newP)
			}
		}

		if solidity > maxSolidity {
			maxSolidity = solidity
		}
	}

	return maxSolidity
}
