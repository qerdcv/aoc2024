package main

import (
	"fmt"
	"os"

	day "github.com/qerdcv/aoc2023/days/day_1"
)

func run() error {
	input, err := os.Open("inputs/day_1.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day.ResolvePartTwo(input))

	return nil
}
