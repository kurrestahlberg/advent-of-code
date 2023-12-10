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
	currents := []string{}

	for i, line := range lines {
		if i == 0 {
			sequence = line
		} else if i == 1 {
			continue
		} else {
			parts := regex.FindStringSubmatch(line)
			network[parts[1]] = Node{parts[2], parts[3]}
			if parts[1][2] == 'A' {
				currents = append(currents, parts[1])
			}
		}
	}

	fmt.Println(currents)

	stepCount := 0
	maxZs := 0
	for {
		for seq := 0; seq < len(sequence); seq++ {
			stepCount += 1
			end := true
			currentZs := 0
			if sequence[seq] == 'R' {
				for i := range currents {
					currents[i] = network[currents[i]].right
					if currents[i][2] != 'Z' {
						end = false
					} else {
						currentZs += 1
					}
				}
			} else {
				for i := range currents {
					currents[i] = network[currents[i]].left
					if currents[i][2] != 'Z' {
						end = false
					} else {
						currentZs += 1
					}
				}
			}
			if end {
				return strconv.Itoa(stepCount)
			}
			if currentZs > maxZs {
				maxZs = currentZs
			}
			if stepCount%1000000 == 0 {
				fmt.Println(stepCount, " -- ", currents, " -- ", maxZs, currentZs)
			}
		}
	}
}
