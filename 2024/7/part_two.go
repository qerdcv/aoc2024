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
	for s.Scan() {
		parts := strings.Split(s.Text(), ": ")
		result, _ := strconv.Atoi(parts[0])
		eqParts := generic.Map(strings.Split(parts[1], " "), func(t string) int {
			i, _ := strconv.Atoi(t)
			return i
		})

		if isValid(result, []string{"+", "*", "||"}, eqParts) {
			total += result
		}
	}

	return total, nil
}
