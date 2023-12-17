package day13

import (
	"bufio"
	"io"
)

type mirror [][]byte

func ResolvePartOne(r io.Reader) int {
	mirrors := parseMirrors(r)
	result := 0
	for _, m := range mirrors {
		result += findMirroring(m, 0)
	}

	return result
}

func ResolvePartTwo(r io.Reader) int {
	mirrors := parseMirrors(r)
	result := 0
	for _, m := range mirrors {
		result += findMirroring(m, 1)
	}

	return result
}

func findMirroring(m mirror, skipCnt int) int {
	return findVerticalMirroring(m, 0, skipCnt) +
		findHorizontalMirroring(m, 0, skipCnt)*100
}

func findVerticalMirroring(m mirror, p1 int, skipCnt int) int {
	rowLen := len(m[0])
	colLen := len(m)
	if p1 >= rowLen {
		return 0
	}

	p2 := p1 + 1
	rIdx := 0
	skipped := 0

	for rIdx < colLen {
		if p2 == rowLen {
			return 0
		}

		if m[rIdx][p1] == m[rIdx][p2] {
			rIdx++
			continue
		} else {
			if skipCnt-skipped > 0 {
				skipped++
				rIdx++
				continue
			}
		}

		skipped = 0
		rIdx = 0
		p1++
		p2++
	}

	// if out of row iteration - possible mirroring
	// need to find lower count of iterations
	iterationCnt := p1
	if p1 > rowLen-p2-1 {
		iterationCnt = rowLen - p2 - 1
	}

	for i := 1; i <= iterationCnt; i++ {
		for _, row := range m {
			if row[p1-i] != row[p2+i] {
				if skipCnt-skipped > 0 {
					skipped++
					continue
				}

				return findVerticalMirroring(m, p2, skipCnt)
			}
		}
	}

	if skipCnt != skipped {
		return findVerticalMirroring(m, p2, skipCnt)
	}

	return p2
}

func findHorizontalMirroring(m mirror, p1 int, skipCnt int) int {
	rowLen := len(m[0])
	colLen := len(m)
	if p1 >= colLen {
		return 0
	}

	p2 := p1 + 1
	cIdx := 0
	skipped := 0

	for cIdx < rowLen {
		if p2 == colLen {
			return 0
		}

		if m[p1][cIdx] == m[p2][cIdx] {
			cIdx++
			continue
		} else {
			if skipCnt-skipped > 0 {
				skipped++
				cIdx++
				continue
			}
		}

		skipped = 0
		cIdx = 0
		p1++
		p2++
	}

	// if out of row iteration - possible mirroring
	// need to find lower count of iterations
	iterationCnt := p1
	if p1 > colLen-p2-1 {
		iterationCnt = colLen - p2 - 1
	}

	for i := 1; i <= iterationCnt; i++ {
		for j := 0; j < rowLen; j++ {
			if m[p1-i][j] != m[p2+i][j] {
				if skipCnt-skipped > 0 {
					skipped++
					continue
				}
				return findHorizontalMirroring(m, p2, skipCnt)
			}
		}
	}

	if skipCnt != skipped {
		return findHorizontalMirroring(m, p2, skipCnt)
	}

	return p2
}

func parseMirrors(r io.Reader) []mirror {
	var mirrors []mirror
	s := bufio.NewScanner(r)

	var m mirror = nil
	for s.Scan() {
		if s.Text() == "" {
			mirrors = append(mirrors, m)
			m = make(mirror, 0)
			continue
		}

		if m == nil {
			m = make(mirror, 0)
		}

		m = append(m, []byte(s.Text()))
	}

	mirrors = append(mirrors, m)
	return mirrors
}
