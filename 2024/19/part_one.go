package main

import (
	"strings"
)

func solvePartOne(input string) (int, error) {
	inputParts := strings.Split(input, "\n\n")

	patterns := strings.Split(inputParts[0], ", ")
	designs := strings.Split(inputParts[1], "\n")

	total := 0
	for _, design := range designs {
		n := len(design)
		dp := make([]bool, n+1)
		dp[0] = true

		for i := 1; i <= n; i++ {
			for _, pattern := range patterns {
				if i >= len(pattern) && design[i-len(pattern):i] == pattern {
					if dp[i-len(pattern)] {
						dp[i] = true
						break
					}
				}
			}
		}

		if dp[n] {
			total++
		}
	}

	return total, nil
}
