package main

import (
	"fmt"
	"os"

	day5 "github.com/qerdcv/aoc2023/days/day_5"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_5.txt")
	assertNoErr(err)

	fmt.Println(day5.ResolvePartTwo(f))
}
