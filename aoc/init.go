package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const mainTemplate = `package main

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

	partOne, err := solvePartOne(f)
	if err != nil {
		fmt.Printf("error part one: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	f.(io.Seeker).Seek(0, io.SeekStart)

	partTwo, err := solvePartTwo(f)
	if err != nil {
		fmt.Printf("error part one: %s\n\n", err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("Results:\n\tPart one: %3d\n\tPart two: %3d\n\n", partOne, partTwo)
}
`

const partFileTemplate = `package main

import "io"

func solvePart%s(r io.Reader) (int, error) {
	return 0, nil
}
`

func initDay(args []string) error {
	if len(args) != 2 {
		return ErrInvalidArguments
	}

	year, day := args[0], args[1]
	if v, err := strconv.Atoi(year); err != nil || v <= 0 {
		return fmt.Errorf("%w: invalid year provided", ErrInvalidArguments)
	}

	if v, err := strconv.Atoi(day); err != nil || v <= 0 {
		return fmt.Errorf("%w: invalid day provided", ErrInvalidArguments)
	}

	if !isExists(year) {
		if err := os.Mkdir(year, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create year directory")
		}
	}

	dayDirPath := path.Join(year, day)
	if !isExists(dayDirPath) {
		if err := os.Mkdir(dayDirPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create year directory")
		}
	}

	if err := initFiles(dayDirPath); err != nil {
		os.RemoveAll(dayDirPath)
		return fmt.Errorf("init files: %w", err)
	}

	return nil
}

func isExists(f string) bool {
	_, err := os.Stat(f)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func initFiles(basePath string) error {
	mainF, err := os.Create(path.Join(basePath, "main.go"))
	if err != nil {
		return fmt.Errorf("os create main.go: %w", err)
	}

	_, err = mainF.Write([]byte(mainTemplate))
	if err != nil {
		return fmt.Errorf("write main content: %w", err)
	}

	if err := mainF.Close(); err != nil {
		return fmt.Errorf("close main file: %w", err)
	}

	for _, part := range []string{"One", "Two"} {
		partF, err := os.Create(path.Join(basePath, "part_"+strings.ToLower(part)+".go"))
		if err != nil {
			return fmt.Errorf("create part %s file: %w", part, err)
		}

		_, err = partF.Write([]byte(fmt.Sprintf(partFileTemplate, part)))
		if err != nil {
			return fmt.Errorf("write part %s content: %w", part, err)
		}

		if err = partF.Close(); err != nil {
			return fmt.Errorf("close part %s file: %w", part, err)
		}
	}

	for _, inpFN := range []string{"input.txt", "input.test.txt"} {
		f, err := os.Create(path.Join(basePath, inpFN))
		if err != nil {
			return fmt.Errorf("create input(%s) file: %w", inpFN, err)
		}

		if err := f.Close(); err != nil {
			return fmt.Errorf("close input(%s) file: %w", inpFN, err)
		}
	}

	return nil
}
