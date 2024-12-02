package main

import (
	"fmt"
	"io"
	"os"

	day "github.com/qerdcv/aoc/days/day_2"
)

func run() error {
	input, err := os.Open("inputs/day_2.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	fmt.Println(day.ResolvePartOne(input))

	input.Seek(0, io.SeekStart)

	fmt.Println(day.ResolvePartTwo(input))

	return nil
}
