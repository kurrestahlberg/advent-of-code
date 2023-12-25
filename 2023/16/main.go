package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

type Beam struct {
	x   int
	y   int
	dir int
}

type Pos struct {
	x int
	y int
}

var dirMap = [4]Pos{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func printWorld(world [][]byte) {
	/*
		for _, line := range world {
			fmt.Println(string(line))
		}
		fmt.Println("==============")
	*/
}

func runBeam(world [][]byte, visited [][]byte, beam Beam) ([][]byte, [][]byte, []Beam) {
	if beam.x < 0 || beam.y < 0 || beam.x >= len(world[0]) || beam.y >= len(world) {
		return world, visited, []Beam{}
	}
	for {
		// fmt.Println("beam", beam)
		visited[beam.y][beam.x] = '#'
		switch world[beam.y][beam.x] {
		case 'X':
			printWorld(visited)
			return world, visited, []Beam{}
		case '\\':
			if beam.dir%2 == 0 {
				beam.dir -= 1
			} else {
				beam.dir += 1
			}
		case '/':
			if beam.dir%2 == 0 {
				beam.dir += 1
			} else {
				beam.dir -= 1
			}
		case '-':
			if beam.dir%2 == 0 {
				world[beam.y][beam.x] = 'X'
				printWorld(visited)
				return world, visited, []Beam{{beam.x - 1, beam.y, 3}, {beam.x + 1, beam.y, 1}}
			}
		case '|':
			if beam.dir%2 == 1 {
				world[beam.y][beam.x] = 'X'
				printWorld(visited)
				return world, visited, []Beam{{beam.x, beam.y - 1, 0}, {beam.x, beam.y + 1, 2}}
			}
		}
		for beam.dir < 0 {
			beam.dir += 4
		}
		for beam.dir > 3 {
			beam.dir -= 4
		}
		beam.x += dirMap[beam.dir].x
		beam.y += dirMap[beam.dir].y
		if beam.x < 0 || beam.y < 0 || beam.x >= len(world[0]) || beam.y >= len(world) {
			printWorld(visited)
			return world, visited, []Beam{}
		}
	}
}

func partOne(filename string) string {
	return strconv.Itoa(solve(filename, Beam{0, 0, 1}))
}
func solve(filename string, startingBeam Beam) int {
	lines := readLinesAsByteArray(filename)
	visited := make([][]byte, len(lines))

	for i := 0; i < len(lines); i++ {
		visited[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(visited[i]); j++ {
			visited[i][j] = '.'
		}
	}

	beams := []Beam{startingBeam}

	for {
		if len(beams) <= 0 {
			break
		}
		beam := beams[0]
		var additionalBeams []Beam
		lines, visited, additionalBeams = runBeam(lines, visited, beam)
		beams = append(beams[1:], additionalBeams...)
	}

	total := 0
	for _, line := range visited {
		for _, tile := range line {
			if tile == '#' {
				total += 1
			}
		}
	}
	return total
}

func partTwo(filename string) string {
	lines := readLinesAsByteArray(filename)

	max := 0

	for i := 0; i < len(lines); i++ {
		val := solve(filename, Beam{0, i, 1})
		if val > max {
			max = val
		}
		val = solve(filename, Beam{len(lines[0]) - 1, i, 3})
		if val > max {
			max = val
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		val := solve(filename, Beam{i, 0, 2})
		if val > max {
			max = val
		}
		val = solve(filename, Beam{i, len(lines) - 1, 0})
		if val > max {
			max = val
		}
	}

	return strconv.Itoa(max)
}
