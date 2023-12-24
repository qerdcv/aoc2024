package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/qerdcv/aoc2023/internal/generic"
)

var ErrDayAlreadyExists = errors.New("day already exists")

const dayFileTemplate = `package day%s

import "io"

func ResolvePartOne(r io.Reader) int {
	return 0
}

func ResolvePartTwo(r io.Reader) int {
	return 0
}
`

func initDay(args []string) error {
	_, day := generic.PopStart(args)

	path := filepath.Join("days", "day_"+day)
	if _, err := os.Stat(path); err == nil {
		return ErrDayAlreadyExists
	}

	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return fmt.Errorf("os mkdir: %w", err)
	}

	if err := Touch(filepath.Join("inputs", "day_"+day+".test.txt")); err != nil {
		return fmt.Errorf("create test input: %w", err)
	}

	if err := Touch(filepath.Join("inputs", "day_"+day+".txt")); err != nil {
		return fmt.Errorf("create input: %w", err)
	}

	f, err := os.Create(filepath.Join(path, "day_"+day+".go"))
	if err != nil {
		return fmt.Errorf("create day file: %w", err)
	}

	defer f.Close()

	if _, err = f.Write([]byte(fmt.Sprintf(dayFileTemplate, day))); err != nil {
		return fmt.Errorf("file write: %w", err)
	}

	return nil
}

func Touch(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("os create: %w", err)
	}

	return f.Close()
}
