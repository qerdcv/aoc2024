package main

import (
	"fmt"
	"os"

	day2 "github.com/qerdcv/aoc2023/days/day_2"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_2.txt")
	assertNoErr(err)

	fmt.Println(day2.ResolvePartTwo(f))
}
