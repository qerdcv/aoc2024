package main

import (
	"fmt"
	"os"

	day9 "github.com/qerdcv/aoc2023/days/day_9"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_9.txt")
	assertNoErr(err)

	fmt.Println(day9.ResolvePartTwo(f))
}
