package main

import (
	"fmt"
	"os"

	day11 "github.com/qerdcv/aoc2023/days/day_11"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_11.txt")
	assertNoErr(err)

	fmt.Println(day11.ResolvePartOne(f))
}
