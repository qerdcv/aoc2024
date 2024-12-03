package main

import (
	"errors"
	"fmt"
	"io"
)

func solvePartTwo(r io.Reader) (int, error) {
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
