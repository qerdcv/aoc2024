package days

import (
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/qerdcv/aoc/internal/xmath"
)

func ResolvePartOne(r io.Reader) (int, error) {
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

func ResolvePartTwo(r io.Reader) (int, error) {
	var leftNums []int
	counter := map[int]int{}

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
		counter[rightNum] += 1
	}

	total := 0
	for _, ln := range leftNums {
		v, ok := counter[ln]
		if !ok {
			continue
		}

		total += ln * v
	}

	return total, nil
}
