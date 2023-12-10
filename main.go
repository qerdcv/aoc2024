package main

import (
	"fmt"
	"os"

	day10 "github.com/qerdcv/aoc2023/days/day_10"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_10.txt")
	assertNoErr(err)

	fmt.Println(day10.ResolvePartTwo(f))
}
