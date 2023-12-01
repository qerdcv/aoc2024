package days

import (
	"bufio"
	"io"
	"unicode"
)

func DayOnePartOne(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	var sum int

	for s.Scan() {
		sum += parseOnlyDigits([]rune(s.Text()))
	}

	return sum, nil
}

func parseOnlyDigits(row []rune) int {
	firstDigit := 0
	lastDigit := 0
	for _, r := range row {
		if !unicode.IsDigit(r) {
			continue
		}

		lastDigit = runeToNum(r)
		if firstDigit == 0 {
			firstDigit = lastDigit * 10
		}
	}

	return firstDigit + lastDigit
}

func DayOnePartTwo(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	var sum int

	for s.Scan() {
		sum += parseAlphaDigits([]rune(s.Text()))
	}

	return sum, nil
}

func parseAlphaDigits(row []rune) int {
	firstDigit := 0
	lastDigit := 0

	rowLen := len(row)
	idx := 0
	for idx != rowLen {
		r := row[idx]

		if unicode.IsDigit(r) {
			idx += 1
			lastDigit = runeToNum(r)

			if firstDigit == 0 {
				firstDigit = lastDigit * 10
			}

			continue
		}

		if idx+2 >= rowLen {
			idx += 1
			continue
		}

		tmpIdx := idx + 2
		digitPref := string([]rune{r, row[idx+1], row[idx+2]})
		// min number len is 3
		switch digitPref {
		case "one":
			lastDigit = 1
		case "two":
			lastDigit = 2
		case "thr":
			if tmpIdx+2 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1], row[tmpIdx+2]}); digit == "three" {
				lastDigit = 3
			}
		case "fou":
			if tmpIdx+1 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1]}); digit == "four" {
				lastDigit = 4
			}
		case "fiv":
			if tmpIdx+1 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1]}); digit == "five" {
				lastDigit = 5
			}
		case "six":
			lastDigit = 6
		case "sev":
			if tmpIdx+2 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1], row[tmpIdx+2]}); digit == "seven" {
				lastDigit = 7
			}
		case "eig":
			if tmpIdx+2 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1], row[tmpIdx+2]}); digit == "eight" {
				lastDigit = 8
			}
		case "nin":
			if tmpIdx+1 >= rowLen {
				break
			}

			if digit := digitPref + string([]rune{row[tmpIdx+1]}); digit == "nine" {
				lastDigit = 9
			}
		}

		if firstDigit == 0 {
			firstDigit = lastDigit * 10
		}

		idx += 1
	}

	return firstDigit + lastDigit
}

func runeToNum(r rune) int {
	return int(r - '0')
}
