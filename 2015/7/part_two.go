package main

import "io"

func solvePartTwo(r io.Reader) (int, error) {
	instructions := parseInstructions(r)
	registers := makeRegisters(instructions)

	registers["b"] = func() uint16 {
		return 46065
	}

	return int(registers["a"]()), nil
}
