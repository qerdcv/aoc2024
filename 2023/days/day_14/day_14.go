package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

const (
	roundedRock byte = 'O'
	cubeRock    byte = '#'
	emptySpace  byte = '.'
)

func ResolvePartOne(r io.Reader) int {
	return calculateLoad(turnNorth(parseMap(r)))
}

func ResolvePartTwo(r io.ReadSeeker) int {
	m := parseMap(r)
	seen := map[string]int{}

	for i := 0; i < 1000000000; i++ {
		m = turnEast(turnSouth(turnWest(turnNorth(m))))
		hash := hashFromMap(m)

		if hashHitCnt, ok := seen[hash]; ok {
			if hashHitCnt > 1 {
				break
			}
			seen[hash] = hashHitCnt + 1
			continue
		}

		seen[hash] = 1
	}

	// calculate cycle len
	cycles := 0
	for _, hit := range seen {
		if hit > 1 {
			cycles += 1
		}
	}

	r.Seek(0, 0)
	m = parseMap(r)
	// seen digit is min number of cycles that needs to be done, to close loop
	seenDigits := int(math.Pow10(int(math.Floor(math.Log10(float64(len(seen)))) + 1)))
	for i := 0; i < 1000000000%(cycles*seenDigits); i++ {
		m = turnEast(turnSouth(turnWest(turnNorth(m))))
	}

	return calculateLoad(m)
}

func turnNorth(m [][]byte) [][]byte {
	// displays end positions of every column
	endPos := make([]int, len(m[0]))

	for rIdx, row := range m {
		for cIdx, c := range row {
			switch c {
			case cubeRock:
				endPos[cIdx] = rIdx + 1
			case emptySpace:
				continue
			case roundedRock:
				if m[endPos[cIdx]][cIdx] == roundedRock {
					endPos[cIdx] += 1
					continue
				}
				m[endPos[cIdx]][cIdx] = roundedRock
				endPos[cIdx] += 1
				m[rIdx][cIdx] = emptySpace
			}
		}
	}

	return m
}

func turnWest(m [][]byte) [][]byte {
	// displays end positions of every column
	endPos := make([]int, len(m))

	for rIdx, row := range m {
		for cIdx, c := range row {
			switch c {
			case cubeRock:
				endPos[rIdx] = cIdx + 1
			case emptySpace:
				continue
			case roundedRock:
				if m[rIdx][endPos[rIdx]] == roundedRock {
					endPos[rIdx] += 1
					continue
				}

				m[rIdx][endPos[rIdx]] = roundedRock
				endPos[rIdx] += 1
				m[rIdx][cIdx] = emptySpace
			}
		}
	}

	return m
}

func turnEast(m [][]byte) [][]byte {
	// displays end positions of every column
	endPos := make([]int, len(m))
	cCnt := len(m[0])
	for idx := range endPos {
		endPos[idx] = cCnt - 1
	}

	for rIdx, row := range m {
		for cIdx := cCnt - 1; cIdx >= 0; cIdx-- {
			c := row[cIdx]
			switch c {
			case cubeRock:
				endPos[rIdx] = cIdx - 1
			case emptySpace:
				continue
			case roundedRock:
				if m[rIdx][endPos[rIdx]] == roundedRock {
					endPos[rIdx] -= 1
					continue
				}

				m[rIdx][endPos[rIdx]] = roundedRock
				endPos[rIdx] -= 1
				m[rIdx][cIdx] = emptySpace
			}
		}
	}

	return m
}

func turnSouth(m [][]byte) [][]byte {
	// displays end positions of every column
	endPos := make([]int, len(m[0]))
	rCnt := len(m)
	for idx := range endPos {
		endPos[idx] = rCnt - 1
	}

	for rIdx := rCnt - 1; rIdx >= 0; rIdx-- {
		row := m[rIdx]
		for cIdx, c := range row {
			switch c {
			case cubeRock:
				endPos[cIdx] = rIdx - 1
			case emptySpace:
				continue
			case roundedRock:
				if m[endPos[cIdx]][cIdx] == roundedRock {
					endPos[cIdx] -= 1
					continue
				}
				m[endPos[cIdx]][cIdx] = roundedRock
				endPos[cIdx] -= 1
				m[rIdx][cIdx] = emptySpace
			}
		}
	}

	return m
}

func printMap(m [][]byte) {
	for _, row := range m {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func hashFromMap(m [][]byte) string {
	hash := ""
	for _, row := range m {
		hash += string(row)
	}

	return hash
}

func calculateLoad(m [][]byte) int {
	rCnt := len(m)
	res := 0
	for rIdx, row := range m {
		for _, c := range row {
			if c == roundedRock {
				res += rCnt - rIdx
			}
		}
	}
	return res
}

func parseMap(r io.Reader) [][]byte {
	s := bufio.NewScanner(r)
	var res [][]byte
	for s.Scan() {
		res = append(res, []byte(s.Text()))
	}

	return res
}
