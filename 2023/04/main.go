package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func partOne() string {
	lines := readLines("input.txt")
	valueExpression := regexp.MustCompile(`Card +(\d+): ([\d ]+) \| ([\d ]+)`)
	numberExpression := regexp.MustCompile(`(\d+)`)
	total := 0
	for _, line := range lines {
		parts := valueExpression.FindStringSubmatch(line)
		if len(parts) == 0 {
			return "ERROR! Invalid line: " + line
		}

		winningNumbersAsStrings := numberExpression.FindAllString(parts[2], -1)
		winningNumbers, _ := sliceAtoi(winningNumbersAsStrings)
		slices.Sort(winningNumbers)
		myNumbersAsStrings := numberExpression.FindAllString(parts[3], -1)
		myNumbers, _ := sliceAtoi(myNumbersAsStrings)
		slices.Sort(myNumbers)

		//fmt.Println(parts[1], "--", winningNumbers, "--", myNumbers)
		myNumbersCounter := 0
		winningNumbersCounter := 0
		lineTotal := 0
		for myNumbersCounter < len(myNumbers) && winningNumbersCounter < len(winningNumbers) {
			if winningNumbers[winningNumbersCounter] == myNumbers[myNumbersCounter] {
				if lineTotal == 0 {
					lineTotal = 1
				} else {
					lineTotal *= 2
				}
				myNumbersCounter += 1
			} else if winningNumbers[winningNumbersCounter] > myNumbers[myNumbersCounter] {
				myNumbersCounter += 1
			} else {
				winningNumbersCounter += 1
			}
		}
		total += lineTotal
	}

	return strconv.Itoa(total)
}

func partTwo() string {
	lines := readLines("input.txt")
	valueExpression := regexp.MustCompile(`Card +(\d+): ([\d ]+) \| ([\d ]+)`)
	numberExpression := regexp.MustCompile(`(\d+)`)
	cards := make([]int, len(lines))
	for i := range cards {
		cards[i] = 1
	}
	total := 0
	for lineNum, line := range lines {
		parts := valueExpression.FindStringSubmatch(line)
		if len(parts) == 0 {
			return "ERROR! Invalid line: " + line
		}

		winningNumbersAsStrings := numberExpression.FindAllString(parts[2], -1)
		winningNumbers, _ := sliceAtoi(winningNumbersAsStrings)
		slices.Sort(winningNumbers)
		myNumbersAsStrings := numberExpression.FindAllString(parts[3], -1)
		myNumbers, _ := sliceAtoi(myNumbersAsStrings)
		slices.Sort(myNumbers)

		// fmt.Println(parts[1], "--", winningNumbers, "--", myNumbers)
		myNumbersCounter := 0
		winningNumbersCounter := 0
		lineTotal := 0
		for myNumbersCounter < len(myNumbers) && winningNumbersCounter < len(winningNumbers) {
			if winningNumbers[winningNumbersCounter] == myNumbers[myNumbersCounter] {
				lineTotal += 1
				myNumbersCounter += 1
			} else if winningNumbers[winningNumbersCounter] > myNumbers[myNumbersCounter] {
				myNumbersCounter += 1
			} else {
				winningNumbersCounter += 1
			}
		}
		// fmt.Println(lineNum, "--", lineTotal)
		for i := lineNum + 1; i <= lineNum+lineTotal; i++ {
			cards[i] += cards[lineNum]
		}
	}

	for _, line := range cards {
		// fmt.Println(lineNum, ":", line)
		total += line
	}

	return strconv.Itoa(total)
}
