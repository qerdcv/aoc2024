package main

import (
	"fmt"
	"os"

	day8 "github.com/qerdcv/aoc2023/days/day_8"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_8.txt")
	assertNoErr(err)

	fmt.Println(day8.ResolvePartTwo(f))
}
