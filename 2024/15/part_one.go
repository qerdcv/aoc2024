package main

import (
	"fmt"
	"strings"

	"github.com/qerdcv/aoc2024/internal/generic"
)

type pos struct {
	x, y int
}

var (
	up    = pos{0, -1}
	down  = pos{0, 1}
	left  = pos{-1, 0}
	right = pos{1, 0}
)

func solvePartOne(input string) (int, error) {
	parts := strings.Split(input, "\n\n")

	grid := generic.Map(strings.Split(parts[0], "\n"), func(t string) []string {
		return strings.Split(t, "")
	})
	moves := generic.Filter(strings.Split(parts[1], ""), func(t string) bool {
		return t != "\n"
	})

	robotPos := findRobotPosition(grid)
	for _, m := range moves {
		dp := moveToDPos(m)
		newRobotPos := pos{
			x: robotPos.x + dp.x,
			y: robotPos.y + dp.y,
		}

		if grid[newRobotPos.y][newRobotPos.x] == "#" {
			continue
		}

		if grid[newRobotPos.y][newRobotPos.x] != "O" {
			grid[robotPos.y][robotPos.x] = "."
			grid[newRobotPos.y][newRobotPos.x] = "@"
			robotPos = newRobotPos
			continue
		}

		start := newRobotPos
		end := start
		isPathFree := false
		for {
			if grid[end.y][end.x] == "." {
				isPathFree = true
				grid[start.y][start.x] = "."
				grid[end.y][end.x] = "O"
				break
			}

			if grid[end.y][end.x] == "#" {
				break
			}

			end = pos{
				x: end.x + dp.x,
				y: end.y + dp.y,
			}
		}

		if isPathFree {
			grid[robotPos.y][robotPos.x] = "."
			grid[newRobotPos.y][newRobotPos.x] = "@"
			robotPos = newRobotPos
			continue
		}
	}

	printGrid(grid)

	total := 0
	for y, row := range grid {
		for x, col := range row {
			if col == "O" {
				total += y*100 + x
			}
		}
	}

	return total, nil
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}

	fmt.Println()
}

func moveToDPos(move string) pos {
	switch move {
	case "<":
		return left
	case "^":
		return up
	case ">":
		return right
	case "v":
		return down
	}

	panic("unknown move: " + move)
}

func findRobotPosition(grid [][]string) pos {
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i][j] == "@" {
				return pos{y: i, x: j}
			}
		}
	}

	return pos{}
}
