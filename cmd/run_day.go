package main

import (
	"fmt"
	"os"

	day15 "github.com/qerdcv/aoc2023/days/day_15"
)

func run() error {
	input, err := os.Open("inputs/day_15.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day15.ResolvePartTwo(input))

	return nil
}
