package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const (
	GamePartLen = 5

	red   = "red"
	green = "green"
	blue  = "blue"
)

var (
	validGamePreconditions = map[string]int{
		red:   12,
		green: 13,
		blue:  14,
	}
)

func ResolvePartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	idSum := 0

	for s.Scan() {
		gameID, isValid := parseValidGame(s.Text())
		if !isValid {
			continue
		}

		idSum += gameID
	}

	return idSum
}

func parseValidGame(line string) (int, bool) {
	// skip "Game" part:
	line = line[GamePartLen:]

	// parse game id:
	gameID := 0
	for {
		d := line[0]
		line = line[1:]
		if d == ':' {
			line = line[1:]
			break
		}

		gameID = gameID*10 + int(d-'0')
	}

	line = strings.Replace(line, ";", ",", -1)

	rounds := strings.Split(line, ", ")

	for _, round := range rounds {
		r := strings.Split(round, " ")
		cnt, _ := strconv.Atoi(r[0])
		color := r[1]
		if cnt > validGamePreconditions[color] {
			return 0, false
		}
	}

	return gameID, true
}

func ResolvePartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	powerSum := 0

	for s.Scan() {
		powerSum += parseGamePower(s.Text())
	}

	return powerSum
}

func parseGamePower(line string) int {
	minPowers := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	line = line[strings.Index(line, ":")+1:]
	line = strings.Replace(line, ";", ",", -1)

	rounds := strings.Split(line, ", ")

	for _, round := range rounds {
		r := strings.Split(round, " ")
		cnt, _ := strconv.Atoi(r[0])
		color := r[1]
		if minPower := minPowers[color]; cnt > minPower {
			minPowers[color] = cnt
		}
	}

	return minPowers[red] * minPowers[green] * minPowers[blue]
}
