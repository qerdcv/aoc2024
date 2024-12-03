package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc/internal/generic"
)

func ResolvePartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	total := 0
	for s.Scan() {
		if isSave(generic.Map[string, int](
			strings.Split(s.Text(), " "),
			func(t string) int {
				i, _ := strconv.Atoi(t)
				return i
			})) {
			total += 1
		}
	}

	return total
}

func isSave(levels []int) bool {
	t := 0
	for i := 0; i+1 < len(levels); i++ {
		if t == 0 {
			if levels[i] > levels[i+1] {
				t = 1 // decreasing
			} else {
				t = -1 // increasing
			}
		}

		diff := (levels[i] - levels[i+1]) * t
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func ResolvePartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	total := 0
	for s.Scan() {
		if isSaveV2(generic.Map[string, int](
			strings.Split(s.Text(), " "),
			func(t string) int {
				i, _ := strconv.Atoi(t)
				return i
			})) {
			total += 1
			continue
		}
	}

	return total
}

func isSaveV2(levels []int) bool {
	tmp := make([]int, len(levels))
	for i := 0; i < len(levels); i++ {
		copy(tmp, levels)

		if isSave(append(tmp[:i], tmp[i+1:]...)) {
			return true
		}
	}

	return false
}
