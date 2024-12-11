package main

import (
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartOne(input string) (int, error) {
	strNums := strings.Split(input, " ")

	total := 0
	for _, st := range strNums {
		stones := []string{st}
		for range 25 {
			for i := 0; i < len(stones); i++ {
				if l := len(stones[i]); l%2 == 0 {
					newNum, _ := strconv.Atoi(stones[i][l/2:])
					newStrNum := strconv.Itoa(newNum)
					if i+1 >= len(stones) {
						stones = append(stones, newStrNum)
					} else {
						stones = generic.Insert(stones, newStrNum, i+1)
					}

					newNum, _ = strconv.Atoi(stones[i][:l/2])
					stones[i] = strconv.Itoa(newNum)
					i++
					continue
				}

				if stones[i] == "0" {
					stones[i] = "1"
					continue
				}

				n, _ := strconv.Atoi(stones[i])
				stones[i] = strconv.Itoa(n * 2024)
			}
		}
		total += len(stones)
	}

	return total, nil
}
