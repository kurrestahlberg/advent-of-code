package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 2 Part 1:", partOne("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test3.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

type Node struct {
	left  string
	right string
}

func partOne(filename string) string {
	lines := readLines(filename)
	var sequence string
	network := map[string]Node{}
	regex := regexp.MustCompile(`^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$`)

	for i, line := range lines {
		if i == 0 {
			sequence = line
		} else if i == 1 {
			continue
		} else {
			parts := regex.FindStringSubmatch(line)
			network[parts[1]] = Node{parts[2], parts[3]}
		}
	}

	current := "AAA"
	stepCount := 0
	for {
		for i := 0; i < len(sequence); i++ {
			stepCount += 1
			if sequence[i] == 'R' {
				current = network[current].right
			} else {
				current = network[current].left
			}
			if current == "ZZZ" {
				return strconv.Itoa(stepCount)
			}
		}
	}
}

func partTwo(filename string) string {
	lines := readLines(filename)
	var sequence string
	network := map[string]Node{}
	regex := regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)
	starts := []string{}

	for i, line := range lines {
		if i == 0 {
			sequence = line
		} else if i == 1 {
			continue
		} else {
			parts := regex.FindStringSubmatch(line)
			network[parts[1]] = Node{parts[2], parts[3]}
			if parts[1][2] == 'A' {
				starts = append(starts, parts[1])
			}
		}
	}

	pathLengths := make([]int, len(starts))

	for startIdx, pos := range starts {
		distanceToEnd := findEnd(pos, sequence, network)
		//fmt.Println(pos, distanceToEnd)
		pathLengths[startIdx] = distanceToEnd
	}

	total := 1
	for _, val := range pathLengths {
		lcm := leastCommonMultiple(total, val)
		// fmt.Println("LCM", total, val, lcm, greatestCommonDivisor(total, val))
		total = lcm
	}
	return strconv.Itoa(total)
}

func findEnd(start string, sequence string, network map[string]Node) int {
	current := start
	stepCount := 0
	for {
		for i := 0; i < len(sequence); i++ {
			stepCount += 1
			if sequence[i] == 'R' {
				current = network[current].right
			} else {
				current = network[current].left
			}

			// fmt.Println(current, i, stepCount)

			if current[2] == 'Z' {
				return stepCount
			}
		}
	}
}

func leastCommonMultiple(a int, b int) int {
	return (a * b) / greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a int, b int) int {
	if b == 0 {
		return a
	}
	return greatestCommonDivisor(b, a%b)
}
