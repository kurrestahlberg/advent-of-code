package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt", 6), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt", 64), "time spent", time.Since(start))
	/*
		start = time.Now()
		fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
		start = time.Now()
		fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
	*/
}

type Position struct {
	x     int
	y     int
	steps int
}

func (p *Position) getId() string {
	return fmt.Sprintf("%03d-%03d-%03d", p.x, p.y, p.steps)
}

func findStart(world []string) Position {
	for y, line := range world {
		for x, c := range line {
			if c == 'S' {
				return Position{x, y, 0}
			}
		}
	}
	return Position{0, 0, 0}
}

var xMap = [...]int{0, 1, 0, -1}
var yMap = [...]int{-1, 0, 1, 0}

func traversePos(world []string, pos Position, visited map[string]int) int {
	id := pos.getId()
	total := 0
	if _, ok := visited[id]; ok {
		return 0
	} else {
		if pos.steps == 0 {
			visited[id] = 0
			return 1
		}

		for i := 0; i < 4; i++ {
			newPos := Position{pos.x + xMap[i], pos.y + yMap[i], pos.steps - 1}
			if newPos.x >= 0 && newPos.x < len(world[0]) && newPos.y >= 0 &&
				newPos.y < len(world) && world[newPos.y][newPos.x] != '#' {
				total += traversePos(world, newPos, visited)
			}
		}
	}
	visited[id] = total
	return total
}

func partOne(filename string, maxSteps int) string {
	visited := make(map[string]int)

	world := readLines(filename)
	start := findStart(world)
	start.steps = maxSteps

	total := traversePos(world, start, visited)

	return strconv.Itoa(total)
}

func partTwo(filename string) string {
	/*
		lines := readLines(filename)

		for _, line := range lines {
		}
	*/
	return "not ready"
}
