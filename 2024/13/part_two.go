package main

import (
	"math"
	"strings"
)

func solvePartTwo(input string) (int, error) {
	behavs := strings.Split(input, "\n\n")

	total := 0
	shift := 10000000000000
	for _, behav := range behavs {
		a, b, p := parseBehav(behav)
		p.x += shift
		p.y += shift

		if tokens := calculateTokens(a, b, p, math.MaxInt); tokens != -1 {
			total += tokens
		}
	}

	return total, nil
}
