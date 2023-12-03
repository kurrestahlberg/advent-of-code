package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func RuneKeys(m map[rune]bool) (keys []string) {
	for k := range m {
		keys = append(keys, string(k))
	}
	return keys
}

func partOne() string {
	lines := readLines("input.txt")
	symbols := make([][]int, len(lines))
	total := 0

	value := regexp.MustCompile(`(\d+)`)

	for lineNum, line := range lines {
		for charNum, c := range line {
			switch c {
			case '@', '-', '*', '/', '$', '%', '=', '+', '&', '#':
				symbols[lineNum] = append(symbols[lineNum], charNum)
			}
		}
	}
	for lineNum, line := range lines {
		indices := value.FindAllStringIndex(line, -1)
		numberStrings := value.FindAllString(line, -1)
		numbers := make([]int, len(numberStrings))
		for i, numberString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numberString)
		}
		for i := -1; i <= 1; i++ {
			if lineNum+i > 0 && lineNum+i < len(lines) {
				numbersCounter := 0
				//for _, symbolPos := range symbols[i+lineNum] {
				for symbolIdx := 0; symbolIdx < len(symbols[i+lineNum]); symbolIdx++ {
					symbolPos := symbols[i+lineNum][symbolIdx]
					if numbersCounter >= len(numbers) {
						break
					}
					for symbolPos > indices[numbersCounter][1] || numbers[numbersCounter] == 0 {
						if numbersCounter+1 >= len(numbers) {
							break
						}
						numbersCounter += 1
					}
					if symbolPos >= indices[numbersCounter][0]-1 && symbolPos <= indices[numbersCounter][1] {
						total += numbers[numbersCounter]
						numbers[numbersCounter] = 0
						numbersCounter += 1
						symbolIdx -= 1 // Ensure we check both sides
					}
				}
			}
		}
		// fmt.Println(lineNum, "-", numbers, total)
	}

	return strconv.Itoa(total)
}

func partTwo() string {
	lines := readLines("input.txt")
	symbols := make([][]int, len(lines))
	counts := make([][]int, len(lines))
	ratios := make([][]int, len(lines))
	total := 0

	value := regexp.MustCompile(`(\d+)`)

	for lineNum, line := range lines {
		for charNum, c := range line {
			switch c {
			case '*':
				symbols[lineNum] = append(symbols[lineNum], charNum)
				counts[lineNum] = append(counts[lineNum], 0)
				ratios[lineNum] = append(ratios[lineNum], 1)
			}
		}
	}
	for lineNum, line := range lines {
		indices := value.FindAllStringIndex(line, -1)
		numberStrings := value.FindAllString(line, -1)
		numbers := make([]int, len(numberStrings))
		for i, numberString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numberString)
		}
		for i := -1; i <= 1; i++ {
			if lineNum+i > 0 && lineNum+i < len(lines) {
				numbersCounter := 0
				//for _, symbolPos := range symbols[i+lineNum] {
				for symbolIdx := 0; symbolIdx < len(symbols[i+lineNum]); symbolIdx++ {
					symbolPos := symbols[i+lineNum][symbolIdx]
					if numbersCounter >= len(numbers) {
						break
					}
					for symbolPos > indices[numbersCounter][1] || numbers[numbersCounter] == 0 {
						if numbersCounter+1 >= len(numbers) {
							break
						}
						numbersCounter += 1
					}
					if symbolPos >= indices[numbersCounter][0]-1 && symbolPos <= indices[numbersCounter][1] {
						counts[i+lineNum][symbolIdx] += 1
						numbersCounter += 1
						symbolIdx -= 1 // Ensure we check both sides
					}
				}
			}
		}
	}
	// fmt.Println(counts)
	for lineNum, line := range lines {
		indices := value.FindAllStringIndex(line, -1)
		numberStrings := value.FindAllString(line, -1)
		numbers := make([]int, len(numberStrings))
		for i, numberString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numberString)
		}
		for i := -1; i <= 1; i++ {
			if lineNum+i > 0 && lineNum+i < len(lines) {
				numbersCounter := 0
				//for _, symbolPos := range symbols[i+lineNum] {
				for symbolIdx := 0; symbolIdx < len(symbols[i+lineNum]); symbolIdx++ {
					if counts[i+lineNum][symbolIdx] != 2 {
						continue
					}
					symbolPos := symbols[i+lineNum][symbolIdx]
					if numbersCounter >= len(numbers) {
						break
					}
					for symbolPos > indices[numbersCounter][1] || numbers[numbersCounter] == 0 {
						if numbersCounter+1 >= len(numbers) {
							break
						}
						numbersCounter += 1
					}
					if symbolPos >= indices[numbersCounter][0]-1 && symbolPos <= indices[numbersCounter][1] {
						ratios[i+lineNum][symbolIdx] *= numbers[numbersCounter]
						numbers[numbersCounter] = 0
						numbersCounter += 1
						symbolIdx -= 1 // Ensure we check both sides
					}
				}
			}
		}
		// fmt.Println(lineNum, "-", numbers, total)
	}
	// fmt.Println(ratios)
	for lineNum, line := range ratios {
		for valueIdx, value := range line {
			if counts[lineNum][valueIdx] == 2 {
				total += value
			}
		}
	}

	return strconv.Itoa(total)
}
