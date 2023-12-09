package day9

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type extrapolator func([][]int) int

func ResolvePartOne(r io.Reader) int {
	histories := parseHistories(r)

	return extrapolateHistories(histories, forwardExtrapolator)
}

func ResolvePartTwo(r io.Reader) int {
	histories := parseHistories(r)
	return extrapolateHistories(histories, backwardExtrapolator)
}

func backwardExtrapolator(layers [][]int) int {
	ll := len(layers)
	val := layers[ll-1][0]
	for i := ll - 2; i >= 0; i-- {
		val = layers[i][0] - val
	}

	return val
}

func forwardExtrapolator(layers [][]int) int {
	ll := len(layers)
	val := layers[ll-1][len(layers[ll-1])-1]
	for i := ll - 2; i >= 0; i-- {
		val += layers[i][len(layers[i])-1]
	}
	return val
}

func extrapolateHistories(histories [][]int, e extrapolator) int {
	sum := 0
	for _, history := range histories {
		sum += extrapolateHistory(history, e)
	}

	return sum
}

func extrapolateHistory(history []int, extrapolator func(layers [][]int) int) int {
	var layers [][]int
	layers = append(layers, history)
	currentLayer := history
	for {
		newLayer := make([]int, 0, len(currentLayer)-1)
		isAllZeros := true
		for i := 1; i < len(currentLayer); i++ {
			val := currentLayer[i] - currentLayer[i-1]
			newLayer = append(newLayer, val)
			if val != 0 {
				isAllZeros = false
			}
		}

		layers = append(layers, newLayer)
		currentLayer = newLayer
		if isAllZeros {
			break
		}
	}

	return extrapolator(layers)
}

func parseHistories(r io.Reader) [][]int {
	s := bufio.NewScanner(r)
	var histories [][]int
	for s.Scan() {
		rawHistory := strings.Split(s.Text(), " ")
		history := make([]int, len(rawHistory))
		for idx, rawH := range rawHistory {
			history[idx], _ = strconv.Atoi(rawH)
		}

		histories = append(histories, history)
	}

	return histories
}
