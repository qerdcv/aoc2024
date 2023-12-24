package main

import (
	"fmt"
	"os"

	day17 "github.com/qerdcv/aoc2023/days/day_17"
)

func run() error {
	input, err := os.Open("inputs/day_17.test.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day17.ResolvePartOne(input))

	return nil
}
