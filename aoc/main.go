package main

import (
	"fmt"
	"os"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func usage() {
	fmt.Printf("Advent Of Code golang template\n")
	fmt.Printf("This program is made to simplify AOC day setup\n\n")
	fmt.Printf("Available commands:\n\n")
	fmt.Printf("\tinit <year> <day>\t create setup for year <year> and day <day>\n")
	fmt.Printf("\thelp\t\t\t display this message\n")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		return
	}

	args, cmd := generic.PopStart(args)
	var err error
	switch cmd {
	case "init":
		err = initDay(args)
	case "help":
		usage()
	default:
		usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
		usage()
		os.Exit(1)
	}
}
