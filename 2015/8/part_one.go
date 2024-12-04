package main

import (
	"bufio"
	"io"
	"strconv"
)

func solvePartOne(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	for s.Scan() {
		t := s.Text()
		str, _ := strconv.Unquote(t)
		total += len(t) - len(str)
	}

	return total, nil
}
