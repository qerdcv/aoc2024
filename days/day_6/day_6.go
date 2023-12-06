package day6

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"unicode"
)

type race struct {
	time     int
	distance int
}

func ResolvePartOne(r io.Reader) int {
	return calculatePossibilities(parseRaces(r))
}

func calculatePossibilities(rcs []race) int {
	result := 1
	for _, rc := range rcs {
		result *= calculatePossibility(rc)
	}

	return result
}

func ResolvePartTwo(r io.Reader) int {
	rc := parseRace(r)

	return calculatePossibility(rc)
}

func calculatePossibility(r race) int {

	minTime := 0
	maxTime := 0
	// find min time
	for i := 1; i <= r.time; i++ {
		if (r.time-i)*i > r.distance {
			minTime = i
			break
		}
	}

	// find max time
	for i := r.time; i > 0; i-- {
		if (r.time-i)*i > r.distance {
			maxTime = i
			break
		}
	}

	return maxTime + 1 - minTime
}

func parseRaces(r io.Reader) []race {
	numRe := regexp.MustCompile(`\d+`)
	var times []int

	s := bufio.NewScanner(r)
	s.Scan()
	for _, rawNum := range numRe.FindAllString(s.Text(), -1) {
		num, _ := strconv.Atoi(rawNum)
		times = append(times, num)
	}

	races := make([]race, len(times))
	s.Scan()
	for idx, rawNum := range numRe.FindAllString(s.Text(), -1) {
		num, _ := strconv.Atoi(rawNum)
		races[idx] = race{
			time:     times[idx],
			distance: num,
		}
	}

	return races
}

func parseRace(r io.Reader) race {
	t := 0

	s := bufio.NewScanner(r)
	s.Scan()

	for _, ch := range s.Text() {
		if unicode.IsDigit(ch) {
			t = t*10 + int(ch-'0')
		}
	}

	s.Scan()
	d := 0
	for _, ch := range s.Text() {
		if unicode.IsDigit(ch) {
			d = d*10 + int(ch-'0')
		}
	}

	return race{
		time:     t,
		distance: d,
	}
}
