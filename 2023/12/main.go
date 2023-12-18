package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//fmt.Println("Test 1 Part 1:", partOne("test1.txt"), "time spent", time.Since(start))
	//start = time.Now()
	fmt.Println("Test 2 Part 1:", partOne("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

func fitSpan(conditions string, span int) bool {
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

var memoList = make(map[string]int)

func fitAll(conditions string, damagedList []int, solutionSoFar string) int {
	key := conditions + fmt.Sprint(damagedList)
	if val, ok := memoList[key]; ok {
		return val
	}

	total := 0
	if len(conditions) == 0 {
		memoList[key] = 0
		return 0
	}
	for i := 0; i < len(conditions)-(damagedList[0]-1); i++ {
		// currentSolution := solutionSoFar
		if fitSpan(conditions[i:], damagedList[0]) {
			// for pos := i; pos < i+damagedList[0]; pos++ {
			// 	currentSolution += "#"
			// }
			if len(damagedList) == 1 {
				if !strings.Contains(conditions[i+damagedList[0]:], "#") {
					//fmt.Println("Solution:", currentSolution+strings.ReplaceAll(conditions[i+damagedList[0]:], "?", "."))
					total += 1
					// } else {
					// fmt.Println("Almost a solution but extra characters left over", currentSolution, conditions[i+damagedList[0]:])
				}
			} else if i+damagedList[0] < len(conditions) {
				// currentSolution += "."
				total += fitAll(conditions[i+damagedList[0]+1:], damagedList[1:], "") //currentSolution)
			}
		}
		// solutionSoFar += "."
		if conditions[i] == '#' {
			break
		}
	}
	memoList[key] = total
	return total
}

func partOne(filename string) string {
	lines := readLines(filename)

	total := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		conditions := parts[0]
		damaged, _ := sliceAtoi(strings.Split(parts[1], ","))
		// fmt.Println("Problem::", conditions, damaged)
		num := fitAll(conditions, damaged, "")
		total += num
		// fmt.Println("After row", i, "total is", total, "latest was", num)
	}

	return strconv.Itoa(total)
}

func partTwo(filename string) string {
	lines := readLines(filename)

	total := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		conditions := parts[0]
		orig_damaged, _ := sliceAtoi(strings.Split(parts[1], ","))
		damaged := orig_damaged
		for repeat := 0; repeat < 4; repeat++ {
			conditions += "?" + parts[0]
			damaged = append(damaged, orig_damaged...)
		}
		// fmt.Println("Problem::", conditions, damaged)
		num := fitAll(conditions, damaged, "")
		total += num
		// fmt.Println("After row", i, "total is", total, "latest was", num)
	}

	return strconv.Itoa(total)
}
