package main

import (
	"fmt"
	"os"

	day6 "github.com/qerdcv/aoc2023/days/day_6"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_6.txt")
	assertNoErr(err)

	fmt.Println(day6.ResolvePartTwo(f))
}
