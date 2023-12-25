package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", solve("test1.txt", false), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", solve("input.txt", false), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", solve("test1.txt", true), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solve("input.txt", true), "time spent", time.Since(start))
}

func checkSingleRowReflection(line string, column int, allowSmudge bool) (bool, bool) {
	columnsToCheck := min(column, len(line)-column)
	smudgeUsed := false
	for k := 0; k < columnsToCheck; k++ {
		if line[column+k] != line[column-k-1] {
			if allowSmudge && !smudgeUsed {
				smudgeUsed = true
			} else {
				return false, false
			}
		}
	}
	return true, smudgeUsed
}

func checkHorizontalReflection(lines []string, column int, allowSmudge bool) bool {
	smudgeUsed := false
	for _, line := range lines {
		ok, smudgeFound := checkSingleRowReflection(line, column, allowSmudge)
		if !ok || (smudgeUsed && smudgeFound) {
			return false
		} else {
			smudgeUsed = smudgeUsed || smudgeFound
		}
	}
	rv := true
	if allowSmudge {
		rv = smudgeUsed
	}
	return rv
}

func checkHorizontalReflections(lines []string, allowSmudge bool) int {
	for i := 1; i < len(lines[0]); i++ {
		if checkHorizontalReflection(lines, i, allowSmudge) {
			return i
		}
	}
	return -1
}

func checkSingleColReflection(lines []string, row int, column int, allowSmudge bool) (bool, bool) {
	rowsToCheck := min(row, len(lines)-row)
	smudgeUsed := false
	for k := 0; k < rowsToCheck; k++ {
		if lines[row+k][column] != lines[row-k-1][column] {
			if allowSmudge && !smudgeUsed {
				smudgeUsed = true
			} else {
				return false, false
			}
		}
	}
	return true, smudgeUsed
}

func checkVerticalReflection(lines []string, row int, allowSmudge bool) bool {
	smudgeUsed := false
	for i := 0; i < len(lines[0]); i++ {
		ok, smudgeFound := checkSingleColReflection(lines, row, i, allowSmudge)
		if !ok || (smudgeUsed && smudgeFound) {
			return false
		} else {
			smudgeUsed = smudgeUsed || smudgeFound
		}
	}
	rv := true
	if allowSmudge {
		rv = smudgeUsed
	}
	return rv
}

func checkVerticalReflections(lines []string, allowSmudge bool) int {
	for i := 1; i < len(lines); i++ {
		ok := checkVerticalReflection(lines, i, allowSmudge)
		if ok {
			return i
		}
	}
	return -1
}

func solve(filename string, allowSmudge bool) string {
	lines := readLines(filename)

	start := 0
	total := 0
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			end := i
			if i == len(lines)-1 {
				end += 1
			}
			// for _, l := range lines[start:end] {
			// 	fmt.Println(l)
			// }
			// fmt.Println("----")
			horizontal := checkHorizontalReflections(lines[start:end], allowSmudge)
			if horizontal > 0 {
				total += horizontal
				//fmt.Println("Horizontal reflection at", horizontal)
			} else {
				vertical := checkVerticalReflections(lines[start:end], allowSmudge)
				if vertical <= 0 {
					fmt.Println("ERROR!!")
					for _, l := range lines[start:end] {
						fmt.Println(l)
					}
					fmt.Println("----")
					os.Exit(666)
				} else {
					//fmt.Println("Vertical reflection at", vertical)
					total += 100 * vertical
				}
			}
			start = i + 1
		}
	}

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
