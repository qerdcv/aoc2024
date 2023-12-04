package main

import (
	"fmt"
	"os"

	day4 "github.com/qerdcv/aoc2023/days/day_4"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_4.txt")
	assertNoErr(err)

	fmt.Println(day4.ResolvePartTwo(f))
}
