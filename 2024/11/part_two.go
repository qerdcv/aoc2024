package main

import (
	"strconv"
	"strings"
)

type cacheKey struct {
	st string
	b  int
}

func solvePartTwo(input string) (int, error) {
	strNums := strings.Split(input, " ")
	cache := map[cacheKey]int{}
	var cached func(string, int) int

	cached = func(st string, b int) (i int) {
		ck := cacheKey{st, b}
		defer func() {
			if _, ok := cache[ck]; !ok {
				cache[ck] = i
			}
		}()

		if v, ok := cache[ck]; ok {
			return v
		}

		if b == 0 {
			i = 1
			return
		}

		if l := len(st); l%2 == 0 {
			newNum, _ := strconv.Atoi(st[l/2:])
			right := strconv.Itoa(newNum)

			newNum, _ = strconv.Atoi(st[:l/2])
			left := strconv.Itoa(newNum)
			i = cached(left, b-1) + cached(right, b-1)
			return
		}

		if st == "0" {
			i = cached("1", b-1)
			return
		}

		n, _ := strconv.Atoi(st)

		i = cached(strconv.Itoa(n*2024), b-1)
		return
	}

	total := 0
	for _, st := range strNums {
		total += cached(st, 75)
	}

	return total, nil
}
