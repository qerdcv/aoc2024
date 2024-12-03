package main

import (
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/qerdcv/aoc/internal/xmath"
)

func solvePartOne(r io.Reader) (int, error) {
	var (
		leftNums  []int
		rightNums []int
	)

	for {
		var (
			leftNum, rightNum int
		)

		_, err := fmt.Fscanf(
			r,
			"%d   %d\n",
			&leftNum,
			&rightNum,
		)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, fmt.Errorf("scanf: %w", err)
		}

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)

	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	total := 0
	for i := range len(leftNums) {
		total += xmath.Abs(leftNums[i] - rightNums[i])
	}

	return total, nil
}
