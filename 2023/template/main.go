package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

func partOne(filename string) string {
	/*
		lines := readLines(filename)

		for _, line := range lines {
		}
	*/
	return "not ready"
}

func partTwo(filename string) string {
	/*
		lines := readLines(filename)

		for _, line := range lines {
		}
	*/
	return "not ready"
}
