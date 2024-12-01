package days

import (
	"container/heap"
	"errors"
	"fmt"
	"io"

	"github.com/qerdcv/aoc2023/internal/generic"
	"github.com/qerdcv/aoc2023/internal/xmath"
)

func ResolvePartOne(r io.Reader) (int, error) {
	var (
		leftPQ  generic.PriorityQueue
		rightPQ generic.PriorityQueue
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
		if err != nil && !errors.Is(err, io.EOF) {
			return 0, fmt.Errorf("scanf: %w", err)
		}

		heap.Push(&leftPQ, leftNum)
		heap.Push(&rightPQ, rightNum)

		if errors.Is(err, io.EOF) {
			break
		}
	}

	total := 0
	for leftPQ.Len() != 0 && rightPQ.Len() != 0 {
		total += xmath.Abs(heap.Pop(&leftPQ).(int) - heap.Pop(&rightPQ).(int))
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
		if err != nil && !errors.Is(err, io.EOF) {
			return 0, fmt.Errorf("scanf: %w", err)
		}

		leftNums = append(leftNums, leftNum)
		counter[rightNum] += 1

		if errors.Is(err, io.EOF) {
			break
		}
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
