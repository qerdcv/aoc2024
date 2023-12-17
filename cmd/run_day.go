package main

import (
	"fmt"
	"os"

	day16 "github.com/qerdcv/aoc2023/days/day_16"
)

func run() error {
	input, err := os.Open("inputs/day_16.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day16.ResolvePartTwo(input))

	return nil
}
