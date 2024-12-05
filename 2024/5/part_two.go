package main

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(r io.Reader) (int, error) {
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
			continue
		}

		fixedUpd := fixUpdate(pgRules, upd)
		total += fixedUpd[len(fixedUpd)/2]
	}

	return total, nil
}

func fixUpdate(rules map[int][]int, upd []int) []int {
	seen := make(map[int]int, len(upd))
	for i, u := range upd {
		seen[u] = i
	}

	for !isValidUpdate(rules, upd) {
		for i, u := range upd {
			if r, ok := rules[u]; ok {
				for k, v := range seen {
					if slices.Contains(r, k) && i > v {
						seen[upd[i]] = v
						upd[v], upd[i] = upd[i], upd[v]
						seen[k] = i
					}
				}
			}
		}
	}

	return upd
}
