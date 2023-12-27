package main

import (
	"fmt"
	"regexp"
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
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

var dirMap = map[string]int{
	"U": 0, //{0, -1},
	"R": 1, //{1, 0},
	"D": 2, //{0, 1},
	"L": 3, //{-1, 0},
}
var dirList = []Coord{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

const defaultMultiplier = 10

func resolveCursorMove(cursorPos Coord, dir int) Coord {
	switch dir {
	case 0:
		return Coord{0 - cursorPos.x, 0}
	case 1:
		return Coord{0, 0 - cursorPos.y}
	case 2:
		return Coord{1 - cursorPos.x, 0}
	case 3:
		return Coord{0, 1 - cursorPos.y}
	}
	return cursorPos
}

func partOne(filename string) string {
	lines := readLines(filename)
	return strconv.Itoa(solve(lines))
}

func solve(lines []string) int {

	//coordinates := make([]Coord, len(lines)+1)
	//coordinates[0] = Coord{0, 0}

	outline := make([]Coord, len(lines)+1)
	outline[0] = Coord{0, 0}

	prevDir := 1
	rightTurns := 0
	leftTurns := 0
	cursorPos := Coord{0, 0}

	for i, line := range lines {
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[1])
		dir := dirMap[parts[0]]
		//coordinates[i+1] = Coord{coordinates[i].x + steps*dirList[dir].x, coordinates[i].y + steps*dirList[dir].y}

		cursorChange := resolveCursorMove(cursorPos, dir)
		outline[i] = Coord{outline[i].x + cursorChange.x, outline[i].y + cursorChange.y}
		cursorPos = Coord{cursorPos.x + cursorChange.x, cursorPos.y + cursorChange.y}

		if (dir > prevDir && (dir != 3 || prevDir != 0)) || (dir == 0 && prevDir == 3) {
			rightTurns += 1
		} else if dir != prevDir {
			leftTurns += 1
		}

		outline[i+1] = Coord{outline[i].x + steps*dirList[dir].x, outline[i].y + steps*dirList[dir].y}
		prevDir = dir
	}

	fmt.Println("Right turns:", rightTurns, "left turns:", leftTurns)

	return shoelace(outline)
}

var regex = regexp.MustCompile(`[URDL] [0-9]+ \(#([0-9a-f]{5})([0-3])\)`)
var numberToDirMap = []string{"R", "D", "L", "U"}

func convertLine(input string) string {
	match := regex.FindStringSubmatch(input)
	val, _ := strconv.ParseUint(match[1], 16, 64)
	dir, _ := strconv.Atoi(match[2])
	return numberToDirMap[dir] + " " + strconv.Itoa(int(val))
}

func partTwo(filename string) string {

	lines := readLines(filename)

	converted := make([]string, len(lines))

	for i, line := range lines {
		converted[i] = convertLine(line)
		//fmt.Println(converted[i])
	}

	return strconv.Itoa(solve(converted))
}
