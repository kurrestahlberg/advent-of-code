package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func partOne() string {
	lines := readLines("input.txt")
	numberExpression := regexp.MustCompile(`(\d+)`)

	times, _ := sliceAtoi(numberExpression.FindAllString(lines[0], -1))
	distances, _ := sliceAtoi(numberExpression.FindAllString(lines[1], -1))

	total := 1

	for timeIndex, time := range times {
		raceTotal := 0
		for i := 1; i < time; i++ {
			distance := (time - i) * i
			if distance > distances[timeIndex] {
				raceTotal += 1
			}
		}
		total *= raceTotal
	}

	return strconv.Itoa(total)
}

func partTwo() string {
	lines := readLines("input.txt")
	numberExpression := regexp.MustCompile(`(\d+)`)

	times, _ := sliceAtoi(numberExpression.FindAllString(strings.Replace(lines[0], " ", "", -1), -1))
	distances, _ := sliceAtoi(numberExpression.FindAllString(strings.Replace(lines[1], " ", "", -1), -1))

	total := 1

	for timeIndex, time := range times {
		raceTotal := 0
		for i := 1; i < time; i++ {
			distance := (time - i) * i
			if distance > distances[timeIndex] {
				raceTotal += 1
			}
		}
		total *= raceTotal
	}

	return strconv.Itoa(total)
}
