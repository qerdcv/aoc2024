package main

import (
	"fmt"
	"os"

	day12 "github.com/qerdcv/aoc2023/days/day_12"
)

func run() error {
	input, err := os.Open("inputs/day_12.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day12.ResolvePartTwo(input))

	return nil
}
