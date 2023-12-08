package day8

import (
	"bufio"
	"io"
	"strings"
)

const (
	firstNode = "AAA"
	lastNode  = "ZZZ"
)

func ResolvePartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	way := parseWay(s)
	nodes := parseNetworkTree(s)
	head := firstNode
	return findWay(nodes, head, way, func(n string) bool {
		return n == lastNode
	})
}

func ResolvePartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	way := parseWay(s)
	nodes, startPositions := parseNetworkTreeWithStartPositions(s)

	return findMultiWay(nodes, startPositions, way, func(n string) bool {
		return n[2] == 'Z'
	})
}

func findMultiWay(
	nodes map[string][2]string,
	startPositions []string,
	way string,
	cond func(n string) bool,
) int {

	totalSteps := make([]int, 0, len(startPositions))

	for _, p := range startPositions {
		steps := 0
	wayLoop:
		for {
			for _, ch := range way {
				steps += 1

				switch ch {
				case 'L':
					p = nodes[p][0]
				case 'R':
					p = nodes[p][1]
				}
			}

			if cond(p) {
				totalSteps = append(totalSteps, steps)
				break wayLoop
			}
		}
	}

	return LCM(totalSteps[0], totalSteps[1], totalSteps[2:]...)
}

func findWay(nodes map[string][2]string, head, way string, cond func(n string) bool) int {
	steps := 0
	for {
		for _, ch := range way {
			steps += 1

			switch ch {
			case 'L':
				head = nodes[head][0]
			case 'R':
				head = nodes[head][1]
			}

			if cond(head) {
				return steps
			}
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parseNetworkTree(s *bufio.Scanner) map[string][2]string {
	nodes := map[string][2]string{}
	for s.Scan() {
		t := s.Text()
		st := strings.Split(t, " = ")

		node := st[0]
		left := st[1][1:4]
		right := st[1][6:9]

		nodes[node] = [2]string{left, right}
	}

	return nodes
}

func parseNetworkTreeWithStartPositions(s *bufio.Scanner) (map[string][2]string, []string) {
	nodes := map[string][2]string{}
	startPositions := make([]string, 0, 6)
	for s.Scan() {
		t := s.Text()
		st := strings.Split(t, " = ")

		node := st[0]
		left := st[1][1:4]
		right := st[1][6:9]

		if node[2] == 'A' {
			startPositions = append(startPositions, node)
		}

		nodes[node] = [2]string{left, right}
	}

	return nodes, startPositions
}

func parseWay(s *bufio.Scanner) string {
	way := ""
	for s.Scan() {
		if s.Text() == "" {
			return way
		}

		way = s.Text()
	}

	return way
}
