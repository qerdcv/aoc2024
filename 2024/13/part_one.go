package main

import (
	"strconv"
	"strings"
)

type pos struct{ x, y int }

func solvePartOne(input string) (int, error) {
	behavs := strings.Split(input, "\n\n")

	total := 0
	for _, behav := range behavs {
		a, b, p := parseBehav(behav)
		if tokens := calculateTokens(a, b, p, 100); tokens != -1 {
			total += tokens
		}
	}

	return total, nil
}

func parseBehav(behav string) (pos, pos, pos) {
	lines := strings.Split(behav, "\n")
	aParts := strings.Split(lines[0], ": ")
	bParts := strings.Split(lines[1], ": ")
	priceParts := strings.Split(lines[2], ": ")

	aCoords := strings.Split(aParts[1], ", ")
	bCoords := strings.Split(bParts[1], ", ")
	pCoords := strings.Split(priceParts[1], ", ")

	a := pos{}
	a.x, _ = strconv.Atoi(strings.Split(aCoords[0], "+")[1])
	a.y, _ = strconv.Atoi(strings.Split(aCoords[1], "+")[1])

	b := pos{}
	b.x, _ = strconv.Atoi(strings.Split(bCoords[0], "+")[1])
	b.y, _ = strconv.Atoi(strings.Split(bCoords[1], "+")[1])

	p := pos{}
	p.x, _ = strconv.Atoi(strings.Split(pCoords[0], "=")[1])
	p.y, _ = strconv.Atoi(strings.Split(pCoords[1], "=")[1])

	return a, b, p
}

func calculateTokens(a, b, p pos, threshold int) int {
	aP := (b.y*p.x - b.x*p.y) / (a.x*b.y - b.x*a.y)
	bP := (-a.y*p.x + a.x*p.y) / (a.x*b.y - b.x*a.y)

	if aP < 0 || aP > threshold || bP < 0 || bP > threshold {
		return -1
	}

	if a.x*aP+b.x*bP != p.x || a.y*aP+b.y*bP != p.y {
		return -1
	}

	return aP*3 + bP
}
