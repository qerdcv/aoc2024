package main

import (
	"io"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(r io.Reader) (int, error) {
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

			maxDist := -1
			nextI := ""
			for c2, d := range distMap[i] {
				if _, ok := visited[c2]; ok {
					continue
				}

				if d > maxDist {
					maxDist = d
					nextI = c2
				}
			}

			if nextI != "" {
				q = append(q, nextI)
				localRes += maxDist
			}
		}

		if localRes > res {
			res = localRes
		}
	}

	return res, nil
}
