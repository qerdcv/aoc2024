package main

import (
	"fmt"
	"io"
	"os"

	day "github.com/qerdcv/aoc/days/day_3"
)

func run() error {
	input, err := os.Open("inputs/day_3.txt")
	if err != nil {
		return fmt.Errorf("os open: %w", err)
	}

	defer input.Close()

	bytes, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("io read all: %w", err)
	}

	fmt.Println(day.ResolvePartOne(bytes))

	input.Seek(0, io.SeekStart)

	fmt.Println(day.ResolvePartTwo(bytes))

	return nil
}
