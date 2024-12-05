package main

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartOne(r io.Reader) (int, error) {
	var updates [][]int
	pgRules := map[int][]int{}
	s := bufio.NewScanner(r)

	for s.Scan() {
		t := s.Text()
		if t == "" {
			break
		}

		rule := generic.Map(strings.Split(t, "|"), func(t string) int {
			i, _ := strconv.Atoi(t)
			return i
		})

		pgRules[rule[0]] = append(pgRules[rule[0]], rule[1])
	}

	for s.Scan() {
		updates = append(updates, generic.Map(strings.Split(s.Text(), ","), func(t string) int {
			i, _ := strconv.Atoi(t)
			return i
		}))
	}

	total := 0
	for _, upd := range updates {
		if isValidUpdate(pgRules, upd) {
			middle := upd[len(upd)/2]
			total += middle
		}
	}

	return total, nil
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	seen := make(map[int]bool, len(update))

	for _, u := range update {
		if r, ok := rules[u]; ok {
			for k := range seen {
				if slices.Contains(r, k) {
					return false
				}
			}
		}

		seen[u] = true
	}

	return true
}
