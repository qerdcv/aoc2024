package day4

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

const linesCnt = 205

type numbersPair struct {
	cardNumbers,
	winningNumbers map[int]bool
}

func ResolvePartOne(r io.Reader) int {
	return getTotalCardValues(parseNumbersTable(r))
}

func getTotalCardValues(numPairs []numbersPair) int {
	totalCardsValue := 0
	for _, numPair := range numPairs {
		cardValue := 0
		for k := range numPair.cardNumbers {
			if numPair.winningNumbers[k] {
				if cardValue == 0 {
					cardValue += 1
					continue
				}

				cardValue += cardValue
			}
		}

		totalCardsValue += cardValue
	}

	return totalCardsValue
}

func ResolvePartTwo(r io.Reader) int {
	return getTotalCardCopies(parseNumbersTable(r))
}

func getTotalCardCopies(numPairs []numbersPair) int {
	totalCardsProcessed := 0

	cardsCopies := map[int]int{}

	for idx, numPair := range numPairs {
		matchCnt := 0

		cardsCopies[idx] += 1

		for k := range numPair.cardNumbers {
			if numPair.winningNumbers[k] {
				matchCnt += 1
			}
		}

		for i := 1; i <= matchCnt; i++ {
			cardsCopies[idx+i] += 1 * cardsCopies[idx]
		}
	}

	for _, cnt := range cardsCopies {
		totalCardsProcessed += cnt
	}

	return totalCardsProcessed
}

func parseNumbersTable(r io.Reader) []numbersPair {
	s := bufio.NewScanner(r)
	res := make([]numbersPair, 0, linesCnt)
	for s.Scan() {
		line := s.Text()
		line = line[strings.Index(line, ":")+2:]
		splittedLine := strings.Split(line, " | ")
		res = append(res, numbersPair{
			cardNumbers:    parseNumbersLine(splittedLine[0]),
			winningNumbers: parseNumbersLine(splittedLine[1]),
		})
	}

	return res
}

func parseNumbersLine(s string) map[int]bool {
	res := map[int]bool{}

	number := 0
	for _, ch := range s {
		if ch == ' ' {
			if number != 0 {
				res[number] = true

				number = 0
			}

			continue
		}

		if unicode.IsDigit(ch) {
			number = number*10 + int(ch-'0')
		}
	}

	res[number] = true

	return res
}
