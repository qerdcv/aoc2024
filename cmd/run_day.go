package main

import (
	"fmt"
	"os"

	day13 "github.com/qerdcv/aoc2023/days/day_13"
)

func run() error {
	input, err := os.Open("inputs/day_13.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day13.ResolvePartTwo(input))

	return nil
}
