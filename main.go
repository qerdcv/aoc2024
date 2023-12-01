package main

import (
	"fmt"
	"os"

	"github.com/qerdcv/aoc2023/days"
)

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day_1.txt")
	assertNoErr(err)

	res, err := days.DayOnePartTwo(f)
	assertNoErr(err)
	fmt.Println("Result:", res)
}
