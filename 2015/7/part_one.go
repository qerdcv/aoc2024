package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseInstructions(r io.Reader) [][]string {
	s := bufio.NewScanner(r)

	var instructions [][]string
	for s.Scan() {
		t := s.Text()
		instructions = append(instructions, strings.Split(t, " "))
	}

	return instructions
}

func solvePartOne(r io.Reader) (int, error) {
	instructions := parseInstructions(r)
	registers := makeRegisters(instructions)

	return int(registers["a"]()), nil
}

func makeRegisters(instructions [][]string) map[string]func() uint16 {
	registers := map[string]func() uint16{}

	for _, parts := range instructions {
		switch parts[0] {
		case "NOT":
			registers[parts[3]] = cached(func() uint16 { return ^registers[parts[1]]() })
		default:
			switch parts[1] {
			case "AND":
				x, err := strconv.Atoi(parts[0])
				if err != nil {
					registers[parts[4]] = cached(func() uint16 { return registers[parts[0]]() & registers[parts[2]]() })
					continue
				}

				registers[parts[4]] = cached(func() uint16 { return uint16(x) & registers[parts[2]]() })
			case "OR":
				registers[parts[4]] = cached(func() uint16 { return registers[parts[0]]() | registers[parts[2]]() })
			case "LSHIFT":
				x, err := strconv.Atoi(parts[2])
				if err != nil {
					panic("non numeric LSHIFT: ")
				}

				registers[parts[4]] = cached(func() uint16 { return registers[parts[0]]() << x })

			case "RSHIFT":
				x, err := strconv.Atoi(parts[2])
				if err != nil {
					panic("non numeric RSHIFT: ")
				}

				registers[parts[4]] = cached(func() uint16 { return registers[parts[0]]() >> x })
			case "->":
				x, err := strconv.Atoi(parts[0])
				if err != nil {
					registers[parts[2]] = cached(func() uint16 { return registers[parts[0]]() })
					continue
				}

				registers[parts[2]] = cached(func() uint16 { return uint16(x) })
			}
		}
	}

	return registers
}

func cached(f func() uint16) func() uint16 {
	cache := uint16(0)
	return func() uint16 {
		if cache != 0 {
			return cache
		}

		cache = f()
		return cache
	}
}
