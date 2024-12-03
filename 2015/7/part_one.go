package main

import (
	"bufio"
	"io"
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
	registers := map[string]uint16{}

	instructions := parseInstructions(r)
	findQueue := []string{"a"}
	for len(findQueue) != 0 {
		w
	}

	return int(registers["a"]), nil
}

//switch parts[0] {
//case "NOT":
//registers[parts[3]] = ^registers[parts[1]]
//default:
//switch parts[1] {
//case "AND":
//x, err := strconv.ParseUint(parts[0], 10, 16)
//if err != nil {
//registers[parts[4]] = registers[parts[0]] & registers[parts[2]]
//continue
//}
//
//registers[parts[4]] = uint16(x) & registers[parts[2]]
//case "OR":
//registers[parts[4]] = registers[parts[0]] | registers[parts[2]]
//case "LSHIFT":
//x, err := strconv.ParseUint(parts[2], 10, 16)
//if err != nil {
//panic("non numeric LSHIFT: " + t)
//}
//
//registers[parts[4]] = registers[parts[0]] << x
//
//case "RSHIFT":
//x, err := strconv.ParseUint(parts[2], 10, 16)
//if err != nil {
//panic("non numeric RSHIFT: " + t)
//}
//
//registers[parts[4]] = registers[parts[0]] >> x
//case "->":
//x, err := strconv.ParseUint(parts[0], 10, 16)
//if err != nil {
//registers[parts[2]] = registers[parts[0]]
//continue
//}
//
//registers[parts[2]] = uint16(x)
//}
//}
