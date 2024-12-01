package day15

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

type sign byte

const (
	dash   sign = '-'
	equals sign = '='
)

type label string

func (l label) hash() int {
	currentValue := 0
	for _, c := range l {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

type element struct {
	label label
	sign  sign
	fl    int
}

func ResolvePartOne(r io.Reader) int {
	sequence := parseSequence(r)
	sum := 0
	for _, s := range sequence {
		currentValue := 0
		for _, c := range s {
			currentValue += int(c)
			currentValue *= 17
			currentValue %= 256
		}

		sum += currentValue
	}

	return sum
}

func ResolvePartTwo(r io.Reader) int {
	sequence := parseSeqStruct(r)
	hm := make(map[int][]element, 255)
	for _, s := range sequence {
		hash := s.label.hash()

		box := hm[hash]
		switch s.sign {
		case dash:
			if len(box) == 0 {
				continue
			}

			rmIdx := slices.IndexFunc(box, func(e element) bool {
				return e.label == s.label
			})
			if rmIdx < 0 {
				continue
			}

			hm[hash] = append(box[:rmIdx], box[rmIdx+1:]...)
			if len(hm[hash]) == 0 {
				delete(hm, hash)
			}

		case equals:
			if len(box) == 0 {
				hm[hash] = append(hm[hash], s)
				continue
			}

			replaceIdx := slices.IndexFunc(box, func(e element) bool {
				return e.label == s.label
			})
			if replaceIdx < 0 {
				hm[hash] = append(hm[hash], s)
				continue
			}

			hm[hash][replaceIdx] = s
		}
	}

	sum := 0
	for boxNum, box := range hm {
		if len(box) == 0 {
			continue
		}

		boxNum += 1
		for idx, e := range box {
			idx += 1
			sum += boxNum * idx * e.fl
		}
	}

	return sum
}

func parseSeqStruct(r io.Reader) []element {
	s := bufio.NewScanner(r)
	s.Scan()
	rawElements := strings.Split(s.Text(), ",")
	elements := make([]element, len(rawElements))
	for idx, el := range rawElements {
		if strings.Contains(el, "-") {
			elements[idx] = element{
				label: label(el[:len(el)-1]),
				sign:  dash,
			}
		} else {
			sepIdx := strings.Index(el, "=")
			elements[idx] = element{
				label: label(el[:sepIdx]),
				sign:  equals,
			}
			elements[idx].fl, _ = strconv.Atoi(el[sepIdx+1:])
		}
	}

	return elements
}

func parseSequence(r io.Reader) []string {
	s := bufio.NewScanner(r)
	s.Scan()
	return strings.Split(s.Text(), ",")
}
