package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func partOne() string {
	first, _ := regexp.Compile(`^\D*(\d)`)
	last, _ := regexp.Compile(`(\d)\D*$`)
	lines := readLines("input.txt")

	sum := 0

	for _, line := range lines {
		match1 := first.FindStringSubmatch(line)
		match2 := last.FindStringSubmatch(line)
		number, _ := strconv.Atoi((match1[1] + match2[1]))
		sum += number
		// fmt.Println(match1[1], "&", match2[1], "=", number, ", sum is", sum)
	}

	return strconv.Itoa(sum)
}

func parseNumber(number string) int {
	switch number {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "zero":
		return 0
	default:
		val, _ := strconv.Atoi(number)
		return val
	}
}

func partTwo() string {
	first, _ := regexp.Compile(`(\d|one|two|three|four|five|six|seven|eight|nine|zero).*`)
	last, _ := regexp.Compile(`.*(\d|one|two|three|four|five|six|seven|eight|nine|zero)`)
	lines := readLines("input.txt")

	sum := 0

	for _, line := range lines {
		match1 := first.FindStringSubmatch(line)
		match2 := last.FindStringSubmatch(line)

		num1 := parseNumber(match1[1])
		num2 := parseNumber(match2[1])

		number := num1*10 + num2
		sum += number
		//fmt.Println(match1[1], "&", match2[1], "=", number, ", sum is", sum)
	}

	return strconv.Itoa(sum)
}
