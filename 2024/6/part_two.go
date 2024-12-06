package main

type pos struct {
	y, x int
}

const (
	up = iota
	right
	down
	left
)

func solvePartTwo(lines [][]byte) (int, error) {
	y, x, dir := getGuardPos(lines)
	secPath := getSecPath(lines, y, x, dir)
	visited := make(map[visitedKey]struct{}, len(secPath))

	total := 0
	for localY := range len(lines) {
		for localX := range len(lines) {
			if ch := lines[localY][localX]; ch == '#' || (localY == y && localX == x) {
				continue
			}

			spd, ok := secPath[pos{localY, localX}]
			if !ok {
				continue
			}

			tmp := lines[localY][localX]
			lines[localY][localX] = '#'

			// optimization to reduce alloc calls
			clear(visited)
			dy, dx := deltaFromDir(spd)
			if isLooped(lines, visited, localY-dy, localX-dx, spd) {
				total += 1
			}

			lines[localY][localX] = tmp
		}
	}

	return total, nil
}

func getSecPath(lines [][]byte, y, x, dir int) map[pos]int {
	res := map[pos]int{}
	for {
		dy, dx := deltaFromDir(dir)
		y += dy
		x += dx
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
			break
		}

		if lines[y][x] == '#' {
			y, x = y-dy, x-dx
			dir = (dir + 1) % 4
			continue
		}

		if _, ok := res[pos{y, x}]; !ok {
			res[pos{y, x}] = dir
		}
	}

	return res
}

type visitedKey struct {
	y, x, dir int
}

func isLooped(lines [][]byte, visited map[visitedKey]struct{}, y, x, dir int) bool {
	dy, dx := deltaFromDir(dir)
	for {
		y, x = y+dy, x+dx
		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
			return false
		}

		if _, ok := visited[visitedKey{y, x, dir}]; ok {
			return true
		}

		if lines[y][x] == '#' {
			y, x = y-dy, x-dx
			dir = (dir + 1) % 4
			dy, dx = deltaFromDir(dir)
			continue
		}

		visited[visitedKey{y, x, dir}] = struct{}{}
	}
}
