package main

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/qerdcv/aoc2024/internal/generic"
)

func solvePartTwo(input string) (int, error) {
	parts := strings.Split(input, "\n\n")

	grid := expandGrid(generic.Map(strings.Split(parts[0], "\n"), func(t string) []string {
		return strings.Split(t, "")
	}))

	// grid := generic.Map(strings.Split(parts[0], "\n"), func(t string) []string {
	// 	return strings.Split(t, "")
	// })
	moves := generic.Filter(strings.Split(parts[1], ""), func(t string) bool {
		return t != "\n"
	})

	fmt.Println(">")
	printGrid(grid)
	time.Sleep(5 * time.Second)

	robotPos := findRobotPosition(grid)
	for n, m := range moves {
		fmt.Println(m)
		printGrid(grid)
		time.Sleep(42 * time.Millisecond)

		dp := moveToDPos(m)
		newRobotPos := pos{
			x: robotPos.x + dp.x,
			y: robotPos.y + dp.y,
		}

		if grid[newRobotPos.y][newRobotPos.x] == "#" {
			continue
		}

		if el := grid[newRobotPos.y][newRobotPos.x]; el != "[" && el != "]" {
			grid[robotPos.y][robotPos.x] = "."
			grid[newRobotPos.y][newRobotPos.x] = "@"
			robotPos = newRobotPos
			continue
		}

		if dp == left || dp == right {
			if !isBoxMovableX(grid, newRobotPos, dp) {
				continue
			}
			current := newRobotPos
			next := pos{x: current.x + dp.x, y: current.y + dp.y}
			tmp := "."
			for {
				tmp, grid[current.y][current.x] = grid[current.y][current.x], tmp
				current = next
				next = pos{x: current.x + dp.x, y: current.y + dp.y}

				if tmp == "." || tmp == "#" {
					break
				}
			}
		} else {
			boxes, ok := isBoxMovableY(grid, newRobotPos, dp)
			if !ok {
				continue
			}

			sort.Slice(boxes, func(i, j int) bool {
				if dp.y > 0 {
					return boxes[i].y > boxes[j].y || boxes[i].x < boxes[j].x
				}

				return boxes[i].y < boxes[j].y || boxes[i].x < boxes[j].x
			})

			for _, bp := range boxes {
				newBP := pos{x: bp.x + dp.x, y: bp.y + dp.y}
				grid[newBP.y][newBP.x], grid[bp.y][bp.x] = grid[bp.y][bp.x], "."
			}
		}

		grid[robotPos.y][robotPos.x] = "."
		grid[newRobotPos.y][newRobotPos.x] = "@"
		robotPos = newRobotPos
		if n == 4 {
			break
		}
	}

	printGrid(grid)

	total := 0
	for y, row := range grid {
		for x, col := range row {
			if col != "[" {
				continue
			}

			total += 100*y + x
		}
	}

	return total, nil
}

// traversal boxes, to find if there"s any obstacle
func isBoxMovableY(grid [][]string, start pos, dp pos) ([]pos, bool) {
	visited := map[pos]bool{}
	q := list.New()
	q.PushBack(start)
	for q.Len() != 0 {
		el := q.Front()
		q.Remove(el)
		p := el.Value.(pos)

		if visited[p] {
			continue
		}

		visited[p] = true
		lp := pos{x: p.x - 1, y: p.y}
		if !visited[lp] && grid[lp.y][lp.x] == "[" {
			q.PushBack(lp)
		}

		rp := pos{x: p.x + 1, y: p.y}
		if !visited[rp] && grid[rp.y][rp.x] == "]" {
			q.PushBack(rp)
		}

		nextP := pos{x: p.x + dp.x, y: p.y + dp.y}
		if visited[nextP] {
			continue
		}

		next := grid[nextP.y][nextP.x]
		if next == "#" {
			return nil, false
		}

		if next == "." {
			continue
		}

		q.PushBack(nextP)
	}

	boxes := make([]pos, 0, len(visited))
	for v := range visited {
		boxes = append(boxes, v)
	}

	return boxes, true
}

func isBoxMovableX(grid [][]string, start pos, dp pos) bool {
	end := pos{x: start.x + dp.x, y: start.y + dp.y}
	for {
		if grid[end.y][end.x] == "#" {
			return false
		}

		if grid[end.y][end.x] == "." {
			return true
		}

		end.x += dp.x
		end.y += dp.y
	}
}

func expandGrid(grid [][]string) [][]string {
	newSize := len(grid) * 2
	newGrid := make([][]string, 0, newSize)
	topNbot := strings.Split(strings.Repeat("#", newSize), "")
	newGrid = append(newGrid, topNbot)
	for i := 1; i < len(grid)-1; i++ {
		line := make([]string, 0, newSize*2)
		for _, el := range grid[i] {
			switch el {
			case "#":
				line = append(line, "#", "#")
			case ".":
				line = append(line, ".", ".")
			case "O":
				line = append(line, "[", "]")
			case "@":
				line = append(line, "@", ".")
			}
		}

		newGrid = append(newGrid, line)
	}

	newGrid = append(newGrid, topNbot)

	return newGrid
}
