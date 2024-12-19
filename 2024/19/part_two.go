package main

import "strings"

func solvePartTwo(input string) (int, error) {
	inputParts := strings.Split(input, "\n\n")

	patterns := strings.Split(inputParts[0], ", ")
	designs := strings.Split(inputParts[1], "\n")

	total := 0
	for _, design := range designs {
		n := len(design)
		dp := make([]int, n+1)
		dp[0] = 1

		for i := 1; i <= n; i++ {
			for _, pattern := range patterns {
				if i >= len(pattern) && design[i-len(pattern):i] == pattern {
					dp[i] += dp[i-len(pattern)]
				}
			}
		}

		total += dp[n]
	}

	return total, nil
}
