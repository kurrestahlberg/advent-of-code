package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Test Part 1:", partOne("test.txt"))
	fmt.Println("Part 1:", partOne("input.txt"))
	fmt.Println("Test Part 2:", partTwo("test.txt"))
	fmt.Println("Part 2:", partTwo("input.txt"))
}

func getNextSequence(numbers []int) []int {
	rv := make([]int, len(numbers)-1)
	for i := 0; i < len(rv); i++ {
		rv[i] = numbers[i+1] - numbers[i]
	}
	return rv
}

func areAllZeroes(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}

func executeStep(numbers []int, appendFirst bool) int {
	sequence := getNextSequence(numbers)
	//fmt.Println("From", numbers, "we got", sequence)
	if areAllZeroes(sequence) {
		//fmt.Println("Sequence", sequence, "causes us to return 0")
		return 0
	}
	newValue := executeStep(sequence, appendFirst)
	if appendFirst {
		//fmt.Println("Sequence", sequence, "causes us to return", sequence[0]-newValue)
		return sequence[0] - newValue
	} else {
		//fmt.Println("Sequence", sequence, "causes us to return", sequence[len(sequence)-1]+newValue)
		return sequence[len(sequence)-1] + newValue
	}
}

func partOne(filename string) string {
	lines := readLines(filename)
	total := 0
	numberExpression := regexp.MustCompile(`(-?\d+)`)

	for _, line := range lines {
		numbers, _ := sliceAtoi(numberExpression.FindAllString(line, -1))
		newValue := executeStep(numbers, false) + numbers[len(numbers)-1]
		//fmt.Println("Original line:", line)
		//fmt.Println("Adding to total", numbers, newValue)
		total += newValue
	}
	return strconv.Itoa(total)
}

func partTwo(filename string) string {
	lines := readLines(filename)
	total := 0
	numberExpression := regexp.MustCompile(`(-?\d+)`)

	for _, line := range lines {
		numbers, _ := sliceAtoi(numberExpression.FindAllString(line, -1))
		newValue := numbers[0] - executeStep(numbers, true)
		//fmt.Println("Original line:", line)
		// fmt.Println("Adding to total", numbers, newValue)
		total += newValue
	}
	return strconv.Itoa(total)
}
