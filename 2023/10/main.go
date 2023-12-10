package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

/*
Directions:

	0 - up
	1 - right
	2 - down
	3 - left

	-1 - invalid
*/

var valid = [4]map[byte]int{
	{
		'F': 1,
		'7': 3,
		'|': 0,
		'L': -1,
		'J': -1,
		'-': -1,
		'.': -1,
	},
	{
		'F': -1,
		'7': 2,
		'|': -1,
		'L': -1,
		'J': 0,
		'-': 1,
		'.': -1,
	},
	{
		'F': -1,
		'7': -1,
		'|': 2,
		'L': 1,
		'J': 3,
		'-': -1,
		'.': -1,
	},
	{
		'F': 2,
		'7': -1,
		'|': -1,
		'L': 0,
		'J': -1,
		'-': 3,
		'.': -1,
	},
}
var moves = [4]Pos{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func checkBounds(lines []string, pos Pos) bool {
	if pos.y < 0 || pos.y > len(lines) || pos.x < 0 || pos.x > len(lines[pos.y]) {
		return false
	}

	return true
}

func findStartingDirection(lines []string, start Pos) int {
	for i, move := range moves {
		y := start.y + move.y
		x := start.x + move.x
		pos := Pos{x, y}

		// fmt.Print("Checking ", pos, ": ")

		// Out of bounds
		if !checkBounds(lines, pos) {
			// fmt.Println("Out of bounds")
			continue
		}
		pipe := pipeAt(lines, pos)
		// fmt.Println(string(pipe), valid[i][pipe])
		if valid[i][pipe] != -1 {
			return i
		}
	}

	// Error!
	fmt.Println("COULD NOT FIND A VALID DIRECTION FROM START!")
	os.Exit(666)
	return -1
}

type Pos struct {
	x int
	y int
}

func move(from Pos, dir int) Pos {
	x := from.x + moves[dir].x
	y := from.y + moves[dir].y
	return Pos{x, y}
}

func pipeAt(lines []string, pos Pos) byte {
	if !checkBounds(lines, pos) {
		fmt.Println("TRYING TO GET PIPE FROM INVALID POS!", pos)
		os.Exit(666)
	}
	return lines[pos.y][pos.x]
}

func partOne(filename string) string {
	lines := readLines(filename)

	pos := Pos{0, 0}
	dir := 0

	pos, dir = findStartPosAndDir(lines, pos, dir)

	// fmt.Println(pos, dir, string(pipeAt(lines, pos)))

	steps := 0
	for {
		steps++
		pos = move(pos, dir)
		dir = valid[dir][pipeAt(lines, pos)]
		if pipeAt(lines, pos) == 'S' {
			break
		}
		// fmt.Println(pos, dir, string(pipeAt(lines, pos)))
	}

	return strconv.Itoa(steps >> 1)
}

func findStartPosAndDir(lines []string, pos Pos, dir int) (Pos, int) {
	for i, line := range lines {
		index := strings.Index(line, "S")
		if index > -1 {
			pos = Pos{index, i}
			dir = findStartingDirection(lines, pos)
		}
	}
	return pos, dir
}

func partTwo(filename string) string {
	lines := readLines(filename)
	mutableLines := make([][]byte, len(lines))
	for i, line := range lines {
		mutableLines[i] = []byte(line)
	}

	pos := Pos{0, 0}
	dir := 0

	pos, dir = findStartPosAndDir(lines, pos, dir)

	// fmt.Println(pos, dir, string(pipeAt(lines, pos)))

	steps := 0
	for {
		steps++
		pos = move(pos, dir)
		dir = valid[dir][pipeAt(lines, pos)]
		pipe := pipeAt(lines, pos)
		if pipe == 'S' {
			break
		}
		if pipe == '|' || pipe == 'J' || pipe == 'L' {
			mutableLines[pos.y][pos.x] = 'X'
		} else {
			mutableLines[pos.y][pos.x] = 'O'
		}
		// fmt.Println(pos, dir, string(pipeAt(lines, pos)))
	}

	count := 0
	inside := false
	for i, line := range mutableLines {
		inside = false
		for j, pipe := range line {
			if pipe == 'X' {
				inside = !inside
			} else if inside && pipe != 'O' && pipe != 'S' {
				mutableLines[i][j] = 'I'
				count += 1
			}
		}
		// fmt.Println("After line", i, "count is", count)
	}

	/*
		for _, line := range mutableLines {
			fmt.Println(string(line))
		}
	*/

	return strconv.Itoa(count)
}
