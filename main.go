package main

import (
	"fmt"
	"os"

	day7 "github.com/qerdcv/aoc2023/days/day_7"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_7.test.txt")
	assertNoErr(err)

	fmt.Println(day7.ResolvePartTwo(f))
}
