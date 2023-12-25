package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math"
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

func moveRocksTo(lines [][]byte, colRange []int, rowRange []int) {
	// fmt.Println("moving rocks", colRange, rowRange)
	for _, row := range rowRange {
		for _, col := range colRange {
			lines[row][col] = 'O'
		}
	}
}

func partOne(filename string) string {
	lines := readLinesAsByteArray(filename)

	moveRocksNorth(lines)

	//fmt.Println(string(line))
	total := calculateWeightOnNorthBeams(lines)
	return strconv.Itoa(total)
}

func calculateWeightOnNorthBeams(lines [][]byte) int {
	total := 0

	for row, line := range lines {
		rowScore := len(lines) - row

		for _, ch := range line {
			if ch == 'O' {
				total += rowScore
			}
		}
	}
	return total
}

func moveRocksNorth(lines [][]byte) {
	for col := 0; col < len(lines[0]); col++ {
		attachedRocks := 0
		for row := len(lines) - 1; row >= 0; row-- {
			switch lines[row][col] {
			case 'O':
				attachedRocks += 1
				lines[row][col] = '.'
			case '#':
				if attachedRocks > 0 {
					moveRocksTo(lines, []int{col}, generateRange(row+1, row+1+attachedRocks))
					attachedRocks = 0
				}
			case '.':
			}
		}
		if attachedRocks > 0 {
			moveRocksTo(lines, []int{col}, generateRange(0, attachedRocks))
		}
	}
}

func moveRocksWest(lines [][]byte) {
	for row := 0; row < len(lines); row++ {
		attachedRocks := 0
		for col := len(lines[0]) - 1; col >= 0; col-- {
			switch lines[row][col] {
			case 'O':
				attachedRocks += 1
				lines[row][col] = '.'
			case '#':
				if attachedRocks > 0 {
					moveRocksTo(lines, generateRange(col+1, col+1+attachedRocks), []int{row})
					attachedRocks = 0
				}
			case '.':
			}
		}
		if attachedRocks > 0 {
			moveRocksTo(lines, generateRange(0, attachedRocks), []int{row})
		}
	}
}

func moveRocksSouth(lines [][]byte) {
	for col := 0; col < len(lines[0]); col++ {
		attachedRocks := 0
		for row := 0; row < len(lines); row++ {
			switch lines[row][col] {
			case 'O':
				attachedRocks += 1
				lines[row][col] = '.'
			case '#':
				if attachedRocks > 0 {
					moveRocksTo(lines, []int{col}, generateRange(row-1, row-1-attachedRocks))
					attachedRocks = 0
				}
			case '.':
			}
		}
		if attachedRocks > 0 {
			moveRocksTo(lines, []int{col}, generateRange(len(lines)-1, len(lines)-1-attachedRocks))
		}
	}
}

func moveRocksEast(lines [][]byte) {
	for row := 0; row < len(lines); row++ {
		attachedRocks := 0
		for col := 0; col < len(lines[0]); col++ {
			switch lines[row][col] {
			case 'O':
				attachedRocks += 1
				lines[row][col] = '.'
			case '#':
				if attachedRocks > 0 {
					moveRocksTo(lines, generateRange(col-1, col-1-attachedRocks), []int{row})
					attachedRocks = 0
				}
			case '.':
			}
		}
		if attachedRocks > 0 {
			moveRocksTo(lines, generateRange(len(lines[0])-1, len(lines[0])-1-attachedRocks), []int{row})
		}
	}
}

func generateRange(start int, end int) []int {
	if end-start == 0 {
		return []int{start}
	}
	step := (end - start) / int(math.Abs(float64(end-start)))
	rv := []int{}
	for i := start; i != end; i += step {
		rv = append(rv, i)
	}
	return rv
}

func partTwo(filename string) string {
	lines := readLinesAsByteArray(filename)

	positionsSeen := make(map[string]int)

	for i := 0; i < 1000000000; i++ {
		runRound(lines)
		hash := getHash(lines)
		if val, ok := positionsSeen[hash]; ok {
			loopLen := i - val
			loopEnd := val + (loopLen * int((1000000000-val)/loopLen))
			roundsNeeded := 1000000000 - loopEnd - 1
			fmt.Println("Hash at", i, "has already been seen at", val)
			fmt.Println("Pre-loop", val, "loop length", loopLen, loopEnd, roundsNeeded)
			for j := 0; j < roundsNeeded; j++ {
				runRound(lines)
			}
			break
		}
		positionsSeen[hash] = i
	}

	return strconv.Itoa(calculateWeightOnNorthBeams(lines))

	// return getHash(lines)
}

func runRound(lines [][]byte) {
	moveRocksNorth(lines)
	moveRocksWest(lines)
	moveRocksSouth(lines)
	moveRocksEast(lines)
}

func getHash(lines [][]byte) string {
	hasher := sha1.New()
	hasher.Write(bytes.Join(lines, []byte{}))
	return hex.EncodeToString(hasher.Sum(nil))
}
