package day10

import (
	"bufio"
	"bytes"
	"io"
)

const (
	verticalPipe   byte = '|'
	horizontalPipe      = '-'
	neConnector         = 'L'
	nwConnector         = 'J'
	swConnector         = '7'
	seConnector         = 'F'
	startPosition       = 'S'
	noPipe              = 0
)

type coord struct {
	x, y int
}

func ResolvePartOne(r io.Reader) int {
	pipeMap := parsePipeMap(r)
	sp := findStartPosition(pipeMap)
	outputPipes := findOutputPipes(sp, pipeMap)
	sp1, sp2 := sp, sp
	p1, p2 := outputPipes[0], outputPipes[1]
	steps := 0
	for {
		steps += 1

		tmpP1 := processPipe(pipeMap, sp1, p1)
		sp1, p1 = p1, tmpP1
		tmpP2 := processPipe(pipeMap, sp2, p2)
		sp2, p2 = p2, tmpP2

		if sp1.y == sp2.y && sp1.x == sp2.x {
			return steps
		}
	}
}

func ResolvePartTwo(r io.Reader) int {
	pipeMap := parsePipeMap(r)
	sp := findStartPosition(pipeMap)
	outputPipes := findOutputPipes(sp, pipeMap)

	sp1, sp2 := sp, sp
	p1, p2 := outputPipes[0], outputPipes[1]

	loopContour := map[coord]bool{
		sp: true,
		p1: true,
		p2: true,
	}

	for {
		tmpP1 := processPipe(pipeMap, sp1, p1)
		sp1, p1 = p1, tmpP1
		tmpP2 := processPipe(pipeMap, sp2, p2)
		sp2, p2 = p2, tmpP2

		loopContour[sp1] = true
		loopContour[sp2] = true

		if sp1.y == sp2.y && sp1.x == sp2.x {
			break
		}
	}

	// need to close the loop, so replace S with corresponding pipe
	pipeMap[sp.y][sp.x] = getMiddlePipe(outputPipes[0], outputPipes[1])
	insidePipeLoopCnt := 0
	for y, row := range pipeMap {
		isInside := false
		var lastPipe byte = noPipe
		for x, ch := range row {
			if !loopContour[coord{x, y}] {
				if isInside {
					insidePipeLoopCnt += 1
				}

				continue
			}

			switch ch {
			case horizontalPipe:
				continue
			case verticalPipe, seConnector, neConnector:
				// if | F L then inside the pipe, if | it can also mean outside the pipe
				isInside = !isInside
				lastPipe = ch
			case nwConnector:
				if lastPipe == neConnector { // IF LJ outside the pipe
					isInside = !isInside
					lastPipe = noPipe
				}
			case swConnector:
				if lastPipe == seConnector { // IF F7 outside the pipe
					isInside = !isInside
					lastPipe = noPipe
				}
			}
		}
	}

	return insidePipeLoopCnt
}

func getMiddlePipe(p1, p2 coord) byte {
	//  0 -2   : |
	// -1  1   : F
	//  1  1   : 7
	// -2  0   : -
	//  1 -1   : J
	// -1 -1   : L

	dx := p1.x - p2.x
	dy := p1.y - p2.y

	if dx == 0 && dy == -2 {
		return verticalPipe
	}

	if dx == -1 && dy == 1 {
		return seConnector
	}

	if dx == 1 && dy == 1 {
		return swConnector
	}

	if dx == -2 && dy == 0 {
		return horizontalPipe
	}

	if dx == 1 && dy == -1 {
		return nwConnector
	}

	if dx == -1 && dy == -1 {
		return neConnector
	}

	return horizontalPipe
}

func processPipe(pm [][]byte, fromPos coord, pipePos coord) coord {
	pipe := pm[pipePos.y][pipePos.x]
	switch pipe {
	case verticalPipe:
		// from top to bottom
		if pipePos.y-1 == fromPos.y {
			return coord{x: pipePos.x, y: pipePos.y + 1}
		}

		// from bottom to top
		if pipePos.y+1 == fromPos.y {
			return coord{x: pipePos.x, y: pipePos.y - 1}
		}
	case horizontalPipe:
		// from left to right
		if pipePos.x-1 == fromPos.x {
			return coord{x: pipePos.x + 1, y: pipePos.y}
		}

		// from right to left
		if pipePos.x+1 == fromPos.x {
			return coord{x: pipePos.x - 1, y: pipePos.y}
		}

	case neConnector:
		// from top to right
		if pipePos.y-1 == fromPos.y {
			return coord{x: pipePos.x + 1, y: pipePos.y}
		}

		// from right to top
		if pipePos.x+1 == fromPos.x {
			return coord{x: pipePos.x, y: pipePos.y - 1}
		}
	case nwConnector:
		// from top to left
		if pipePos.y-1 == fromPos.y {
			return coord{x: pipePos.x - 1, y: pipePos.y}
		}

		// from left to top
		if pipePos.x-1 == fromPos.x {
			return coord{x: pipePos.x, y: pipePos.y - 1}
		}
	case swConnector:
		// from left to bottom
		if pipePos.x-1 == fromPos.x {
			return coord{x: pipePos.x, y: pipePos.y + 1}
		}

		// from bottom to left
		if pipePos.y+1 == fromPos.y {
			return coord{x: pipePos.x - 1, y: pipePos.y}
		}

	case seConnector:
		// from right to bottom
		if pipePos.x+1 == fromPos.x {
			return coord{x: pipePos.x, y: pipePos.y + 1}
		}

		// from bottom to right
		if pipePos.y+1 == fromPos.y {
			return coord{x: pipePos.x + 1, y: pipePos.y}
		}

	}

	return coord{fromPos.x, fromPos.y}
}

func findOutputPipes(sp coord, pm [][]byte) [2]coord {
	pipes := [2]coord{}
	pipeIdx := 0
	// check top
	if y := sp.y - 1; y >= 0 && bytes.ContainsRune([]byte{verticalPipe, swConnector, seConnector}, rune(pm[y][sp.x])) {
		pipes[pipeIdx] = coord{y: y, x: sp.x}
		pipeIdx += 1
	}

	// check bot
	if y := sp.y + 1; y < len(pm) && bytes.ContainsRune([]byte{verticalPipe, nwConnector, neConnector}, rune(pm[y][sp.x])) {
		pipes[pipeIdx] = coord{y: y, x: sp.x}
		pipeIdx += 1
	}

	// check left
	if x := sp.x - 1; x >= 0 && bytes.ContainsRune([]byte{horizontalPipe, seConnector, neConnector}, rune(pm[sp.y][x])) {
		pipes[pipeIdx] = coord{y: sp.y, x: x}
		pipeIdx += 1
	}

	// check right
	if x := sp.x + 1; x < len(pm[0]) && bytes.ContainsRune([]byte{horizontalPipe, nwConnector, swConnector}, rune(pm[sp.y][x])) {
		pipes[pipeIdx] = coord{y: sp.y, x: x}
		pipeIdx += 1
	}

	return pipes
}

func findStartPosition(pm [][]byte) coord {
	for y, row := range pm {
		for x, col := range row {
			if col == startPosition {
				return coord{x, y}
			}
		}
	}

	return coord{}
}

func parsePipeMap(r io.Reader) [][]byte {
	s := bufio.NewScanner(r)
	var result [][]byte
	for s.Scan() {
		result = append(result, []byte(s.Text()))
	}

	return result
}
