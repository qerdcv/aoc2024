package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartOne(r io.Reader) (int, error) {
	distMap := parseDistanceMap(r)
	res := -1
	for c := range distMap {
		visited := map[string]struct{}{}
		q := []string{c}
		var i string
		localRes := 0
		for len(q) != 0 {
			q, i = generic.PopStart(q)

			visited[i] = struct{}{}

			minDist := -1
			nextI := ""
			for c2, d := range distMap[i] {
				if _, ok := visited[c2]; ok {
					continue
				}

				if minDist == -1 || d < minDist {
					minDist = d
					nextI = c2
				}
			}

			if nextI != "" {
				q = append(q, nextI)
				localRes += minDist
			}
		}

		if res == -1 || localRes < res {
			res = localRes
		}
	}

	return res, nil
}

func parseDistanceMap(r io.Reader) map[string]map[string]int {
	res := make(map[string]map[string]int)

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " = ")
		dist, _ := strconv.Atoi(parts[1])
		cityParts := strings.Split(parts[0], " to ")
		if _, ok := res[cityParts[0]]; !ok {
			res[cityParts[0]] = map[string]int{}
		}

		if _, ok := res[cityParts[1]]; !ok {
			res[cityParts[1]] = map[string]int{}
		}

		res[cityParts[0]][cityParts[1]] = dist
		res[cityParts[1]][cityParts[0]] = dist
	}

	return res
}
