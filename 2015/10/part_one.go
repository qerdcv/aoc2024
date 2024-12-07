package main

import (
	"io"
	"math"
	"strconv"
)

func solvePartOne(r io.Reader) (int, error) {
	b, _ := io.ReadAll(r)
	res := float64(len(string(b)))

	return int(math.Ceil(res * math.Pow(1.3036, 40))), nil
}

func lookAndSay(inp string) string {
	res := ""
	cnt := 0
	prev := byte(' ')
	for i := len(inp) - 1; i >= 0; i-- {
		if prev == ' ' {
			prev = inp[i]
		}

		current := inp[i]

		if current != prev {
			res = strconv.Itoa(cnt) + string(prev) + res
			prev = current
			cnt = 1
		} else {
			cnt++
		}
	}

	res = strconv.Itoa(cnt) + string(prev) + res

	return res
}
