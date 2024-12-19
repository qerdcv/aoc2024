package main

import (
	"embed"
	"fmt"
	"io"
	"os"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	inputFN := "input.txt"
	if len(os.Args) != 1 && os.Args[1] == "-t" {
		inputFN = "input.test.txt"
	}

	f, err := inputs.Open(inputFN)
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	input := string(b)
	partOne, err := solvePartOne(input)
	if err != nil {
		fmt.Printf("error part one: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	partTwo, err := solvePartTwo(input)
	if err != nil {
		fmt.Printf("error part one: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("Results:\n\tPart one: %3d\n\tPart two: %3d\n\n", partOne, partTwo)
}