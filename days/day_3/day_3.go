package day3

import (
	"bufio"
	"fmt"
	"io"
)

var (
	colModifiers = []int{-1, 0, 1}
	rowModifiers = []int{-1, 0, 1}
)

func ResolvePartOne(r io.Reader) int {
	partsMap := parseMap(r)
	return searchParts(partsMap)
}

func searchParts(data [][]byte) int {
	rowCnt := len(data)
	colCnt := len(data[0])

	digit := 0
	isPart := false
	sum := 0

	for rowIdx, row := range data {
		for colIdx, col := range row {
			if isDigit(col) { // is digit
				digit = digit*10 + int(col-'0')
				// check near of the digit, is it a part

				if !isPart {
				outerLoop:
					for _, rowModifier := range rowModifiers {
						for _, colModifier := range colModifiers {
							if rowModifier == 0 && colModifier == 0 {
								continue
							}

							tmpColIdx := colIdx + colModifier
							tmpRowIdx := rowIdx + rowModifier

							if tmpColIdx < 0 || tmpRowIdx < 0 || tmpColIdx >= colCnt || tmpRowIdx >= rowCnt {
								continue
							}

							chToCheck := data[tmpRowIdx][tmpColIdx]
							if isDigit(chToCheck) || chToCheck == '.' {
								continue
							}

							isPart = true
							break outerLoop
						}
					}
				}

				continue
			}

			if isPart {
				sum += digit
			}

			isPart = false
			digit = 0
		}
	}

	return sum
}

func ResolvePartTwo(r io.Reader) int {
	return searchGearsTotalRatio(parseMap(r))
}

func searchGearsTotalRatio(data [][]byte) int {
	rowCnt := len(data)
	colCnt := len(data[0])

	sum := 0
	nums := [2]int{0, 0}

gearLoop:
	for rowIdx, row := range data {
		for colIdx, col := range row {
			if col != '*' { // we interested only in * char
				continue
			}

			// for specific gear we need visited indexes
			visited := map[string]bool{}
			numCnt := 0

			// we need to search numbers near of the gear
			for _, rowModifier := range rowModifiers {
				for _, colModifier := range colModifiers {
					// skip gear
					if rowModifier == 0 && colModifier == 0 {
						continue
					}

					tmpColIdx := colIdx + colModifier
					tmpRowIdx := rowIdx + rowModifier

					if tmpColIdx < 0 || tmpRowIdx < 0 || tmpColIdx >= colCnt || tmpRowIdx >= rowCnt {
						continue
					}

					chToCheck := data[tmpRowIdx][tmpColIdx]
					if !isDigit(chToCheck) { // we need to check only digits
						continue
					}

					// move left until found the beginning of the number
					for {

						if isDigit(data[tmpRowIdx][tmpColIdx]) {
							if tmpColIdx == 0 {
								break
							}

							tmpColIdx -= 1
							continue
						}

						tmpColIdx += 1
						break
					}

					visitedIdx := ""
					number := 0
					for { // parse found digit
						dig := data[tmpRowIdx][tmpColIdx]
						if !isDigit(dig) {
							break
						}

						visitedIdx += fmt.Sprint(tmpRowIdx, tmpColIdx)
						number = number*10 + int(dig-'0')
						tmpColIdx += 1
						if tmpColIdx >= colCnt {
							break
						}
					}

					// skip visited numbers
					if visited[visitedIdx] {
						continue
					}

					numCnt += 1

					if numCnt > 2 {
						continue gearLoop
					}

					nums[numCnt-1] = number
					visited[visitedIdx] = true
				}
			}

			if numCnt != 2 {
				continue
			}

			sum += nums[0] * nums[1]
		}
	}

	return sum
}

func parseMap(r io.Reader) [][]byte {
	res := make([][]byte, 0, 140)
	s := bufio.NewScanner(r)
	for s.Scan() {
		res = append(res, []byte(s.Text()))
	}

	return res
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
