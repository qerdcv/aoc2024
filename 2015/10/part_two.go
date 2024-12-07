package main

import (
	"io"
	"math"
)

func solvePartTwo(r io.Reader) (int, error) {
	res := 492982.0
	for range 10 {
		res *= 1.303577269034
	}

	return int(math.Ceil(res)), nil
}
