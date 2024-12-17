package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(input string) (int, error) {
	parts := strings.Split(input, "\n\n")
	registers := parseRegisters(parts[0])
	rawProgram := strings.Split(parts[1], ": ")[1]
	program := generic.Map(strings.Split(
		rawProgram,
		",",
	), func(t string) int {
		i, _ := strconv.Atoi(t)
		return i
	})

	rawProgram = rawProgram + ","
	stdout := &strings.Builder{}

	// 2,4,1,5,7,5,4,3,1,6,0,3,5,5,3,0
	// 105734774296329
	b := 105734774296328
	smallest := b
	for i := b; ; i-- {
		registers["A"] = i
		executeProgram(stdout, registers, program)

		prog := stdout.String()
		if prog == rawProgram {
			if i < smallest {
				smallest = i
				fmt.Println(smallest)
			}
		}

		if strings.Contains(prog, "2,4,1,5,7,5,4,3,1,6,0,3,5,5,3") {

			fmt.Printf("dec(%d) oct(%o) res(%s)\n", i, i, prog[:len(prog)-1])
		}
		stdout.Reset()
	}

	return 0, nil
}
