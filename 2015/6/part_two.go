package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	greed := make([][]int, 1000)
	for i := range greed {
		greed[i] = make([]int, 1000)
	}

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		cmd := parts[0]
		if cmd == "turn" {
			cmd += parts[1]
			parts = append(parts[:1], parts[2:]...)
		}

		startParts := strings.Split(parts[1], ",")
		stopParts := strings.Split(parts[3], ",")

		startY, _ := strconv.Atoi(startParts[0])
		startX, _ := strconv.Atoi(startParts[1])

		stopY, _ := strconv.Atoi(stopParts[0])
		stopX, _ := strconv.Atoi(stopParts[1])

		for y := startY; y <= stopY; y++ {
			for x := startX; x <= stopX; x++ {
				switch cmd {
				case "turnon":
					greed[y][x] += 1
				case "turnoff":
					if greed[y][x] > 0 {
						greed[y][x] -= 1
					}
				case "toggle":
					greed[y][x] += 2
				default:
					panic("unknown command: " + cmd)
				}
			}
		}
	}

	generic.ForEach(greed, func(t []int) {
		generic.ForEach(t, func(v int) {
			total += v
		})
	})

	return total, nil
}
