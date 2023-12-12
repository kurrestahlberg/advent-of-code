package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", solve("test.txt", 1), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", solve("input.txt", 1), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 1 Part 2:", solve("test.txt", 0), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 2 Part 2:", solve("test.txt", 9), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 3 Part 2:", solve("test.txt", 99), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solve("input.txt", 999999), "time spent", time.Since(start))
}

type Pos struct {
	row int
	col int
}

func solve(filename string, expansion int) string {
	lines := readLines(filename)

	galaxiesOnRows := make([]int, len(lines))
	galaxiesOnCols := make([]int, len(lines[0]))
	galaxies := []Pos{}

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Pos{i, j})
				galaxiesOnCols[j] += 1
				galaxiesOnRows[i] += 1
			}
		}
	}

	mapExpansion(galaxiesOnCols, expansion)
	mapExpansion(galaxiesOnRows, expansion)

	total := 0

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy2 := galaxies[j]
			coldistance := math.Abs(float64((galaxy.col + galaxiesOnCols[galaxy.col]) - (galaxy2.col + galaxiesOnCols[galaxy2.col])))
			rowdistance := math.Abs(float64((galaxy.row + galaxiesOnRows[galaxy.row]) - (galaxy2.row + galaxiesOnRows[galaxy2.row])))
			total += int(coldistance) + int(rowdistance)
		}
	}

	//fmt.Println(galaxies)
	//fmt.Println(galaxiesOnCols)
	//fmt.Println(galaxiesOnRows)
	return strconv.Itoa(total)
}

func mapExpansion(galaxiesOnCols []int, expansion int) {
	for i := range galaxiesOnCols {
		prev := 0
		if i > 0 {
			prev = galaxiesOnCols[i-1]
		}
		if galaxiesOnCols[i] == 0 {
			galaxiesOnCols[i] = prev + expansion
		} else {
			galaxiesOnCols[i] = prev
		}
	}
}

