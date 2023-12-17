package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test 1 Part 1:", partOne("test1.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 2 Part 1:", partOne("test2.txt"), "time spent", time.Since(start))
/*
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
	*/
}

func fitSpan(conditions string, span int) bool {
	//fmt.Print("Testing fit ", span, " starting from ", conditions)
	for j := 0; j < span; j++ {
		if conditions[j] == '.' {
			//fmt.Println(" -- Can't fit the whole span", j, conditions)
			return false
		}
	}
	// Ensure there's a working spring after the span of failing ones
	if span < len(conditions) {
		if conditions[span] == '#' {
			//fmt.Println(" -- No working spring after span")
			return false
		}
	}
	//fmt.Println(" -- Success")
	return true
}

func fitAll(conditions string, damagedList []int) int {
	total := 0
	if len(conditions) == 0{
		return 0
	}
	//fmt.Println("Testing", damagedList[0], "starting from", startPos, damagedList, conditions[startPos:])
	for i := 0; i < len(conditions) - (damagedList[0] - 1); i++ {
		//fmt.Println(i, startPos, len(conditions), damagedList[0])
		if fitSpan(conditions[i:], damagedList[0]) {
			if len(damagedList) == 1 {
				//fmt.Println("Add one!")
				total += 1
			} else if i + damagedList[0] < len(conditions) {
				total += fitAll(conditions[i + damagedList[0] + 1:], damagedList[1:])
			}
		}
		if conditions[i] == '#' {
			break
		}
	}
	return total
}

func partOne(filename string) string {
	lines := readLines(filename)

	total := 0

	for i, line := range lines {
		parts := strings.Split(line, " ")
		conditions := parts[0]
		damaged, _ := sliceAtoi(strings.Split(parts[1], ","))
		num := fitAll(conditions, damaged)
		total += num
		fmt.Println("After row", i, "total is", total, "latest was", num)
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
