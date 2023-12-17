package main

import (
	"fmt"
	"os"

	day14 "github.com/qerdcv/aoc2023/days/day_14"
)

func run() error {
	input, err := os.Open("inputs/day_14.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day14.ResolvePartTwo(input))

	return nil
}
