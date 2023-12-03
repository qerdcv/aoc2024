package main

import (
	"fmt"
	"os"

	day3 "github.com/qerdcv/aoc2023/days/day_3"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_3.txt")
	assertNoErr(err)

	fmt.Println(day3.ResolvePartTwo(f))
}
