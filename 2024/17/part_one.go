package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartOne(input string) (int, error) {
	parts := strings.Split(input, "\n\n")
	registers := parseRegisters(parts[0])
	program := generic.Map(strings.Split(
		strings.Split(parts[1], ": ")[1],
		",",
	), func(t string) int {
		i, _ := strconv.Atoi(t)
		return i
	})

	stdout := &strings.Builder{}
	executeProgram(stdout, registers, program)
	fmt.Println(stdout.String(), registers)
	return 0, nil
}

func executeProgram(stdout io.Writer, registers map[string]int, program []int) {
	programLen := len(program)
	comboOp := func(i int) int {
		switch i {
		case 0, 1, 2, 3:
			return i
		case 4:
			return registers["A"]
		case 5:
			return registers["B"]
		case 6:
			return registers["C"]
		default:
			panic("invalid compbo op: " + strconv.Itoa(i))
		}
	}

	pc := 0
	for pc < programLen {
		opCode := program[pc]
		pc++
		switch opCode {
		case 0: // adv
			operand := comboOp(program[pc])
			registers["A"] /= 1 << operand
		case 1: // bxl
			operand := program[pc]
			registers["B"] ^= operand
		case 2: // bst
			operand := comboOp(program[pc])
			registers["B"] = operand % 8
		case 3: // jnz
			operand := program[pc]
			if registers["A"] != 0 {
				pc = operand
				continue
			}
		case 4: // bxc
			registers["B"] ^= registers["C"]
		case 5: // out
			operand := comboOp(program[pc])
			fmt.Fprintf(stdout, "%d,", operand%8)
		case 6: // bdv
			operand := comboOp(program[pc])
			registers["B"] = registers["A"] / (1 << operand)
		case 7: // cdv
			operand := comboOp(program[pc])
			registers["C"] = registers["A"] / (1 << operand)
		}

		pc++
	}
}

func parseRegisters(input string) map[string]int {
	rawRegisters := strings.Split(input, "\n")
	registers := make(map[string]int, len(rawRegisters))
	for _, line := range rawRegisters {
		rParts := strings.Split(line, ": ")
		registers[rParts[0][9:]], _ = strconv.Atoi(rParts[1])
	}

	return registers
}
