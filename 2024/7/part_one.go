package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartOne(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	for s.Scan() {
		parts := strings.Split(s.Text(), ": ")
		result, _ := strconv.Atoi(parts[0])
		eqParts := generic.Map(strings.Split(parts[1], " "), func(t string) int {
			i, _ := strconv.Atoi(t)
			return i
		})

		if isValid(result, []string{"+", "*"}, eqParts) {
			total += result
		}
	}

	return total, nil
}

func calculate(parts []int, operators []string) int {
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		num := parts[i]
		operator := operators[i-1]
		if operator == "+" {
			result += num
		} else if operator == "*" {
			result *= num
		} else if operator == "||" {
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(num))
		}
	}
	return result
}

func generateCombinations(ops []string, cnt int) [][]string {
	var generateOperators func(int, []string)
	var combinations [][]string

	generateOperators = func(index int, combination []string) {
		if index == cnt {
			cc := make([]string, cnt)
			copy(cc, combination)
			combinations = append(combinations, cc)
			return
		}

		for _, op := range ops {
			combination[index] = op
			generateOperators(index+1, combination)
		}
	}
	generateOperators(0, make([]string, cnt))
	return combinations
}

func isValid(res int, ops []string, parts []int) bool {
	for _, op := range generateCombinations(ops, len(parts)-1) {
		if calculate(parts, op) == res {
			return true
		}
	}

	return false
}
