package main

import (
	"bufio"
	"io"
	"strconv"
)

func solvePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	for s.Scan() {
		t := s.Text()
		total += len(strconv.Quote(t)) - len(t)
	}

	return total, nil
}
