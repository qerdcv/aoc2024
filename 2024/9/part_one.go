package main

import (
	"io"
	"strings"
)

func solvePartOne(r io.Reader) (int, error) {
	b, _ := io.ReadAll(r)
	input := string(b)

	newLine := ""
	id := 0
	var idMap []int
	for i := 0; i < len(input); i += 2 {
		blockSize := int(input[i] - '0')
		freeSize := 0
		if i+1 < len(input) {
			freeSize = int(input[i+1] - '0')
		}

		newLine += strings.Repeat("1", blockSize) + strings.Repeat(".", freeSize)

		localIDMap := make([]int, blockSize+freeSize)
		for i := range localIDMap {
			localIDMap[i] = id
		}

		idMap = append(idMap, localIDMap...)
		id++
	}

	newLineBytes := []byte(newLine)
	for i, j := 0, len(newLineBytes)-1; i < len(newLineBytes) || j >= 0; {
		if i > j {
			newLineBytes = newLineBytes[:i]
			break
		}

		if newLineBytes[i] != '.' {
			i++
			continue
		}

		if newLineBytes[j] == '.' {
			j--
			continue
		}

		newLineBytes[j], newLineBytes[i] = newLineBytes[i], newLineBytes[j]
		idMap[i], idMap[j] = idMap[j], idMap[i]
	}

	result := 0
	for i := range newLineBytes {
		result += i * idMap[i]
		id += 1
	}

	return result, nil
}
