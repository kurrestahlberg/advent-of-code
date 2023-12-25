package main

import (
	"fmt"
	"strconv"
	"strings"
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

func hash(input string) int {
	current := 0
	for i := 0; i < len(input); i++ {
		current += int(byte(input[i]))
		current *= 17
		current %= 256
	}
	return current
}

func partOne(filename string) string {
	lines := readLines(filename)
	input := strings.Split(strings.Join(lines, ","), ",")
	total := 0
	for _, item := range input {
		total += hash(item)
	}

	return strconv.Itoa(total)
}

type HashMapItem struct {
	label       string
	focalLength int
}

func partTwo(filename string) string {
	lines := readLines(filename)
	input := strings.Split(strings.Join(lines, ","), ",")

	hashMap := [256][]HashMapItem{}

	for _, inputItem := range input {
		if inputItem[len(inputItem)-1] == '-' {
			label := inputItem[:len(inputItem)-1]
			h := hash(label)
			// fmt.Println("Removing", label, "from bucket", h)
			for index, item := range hashMap[h] {
				if item.label == label {
					hashMap[h] = append(hashMap[h][:index], hashMap[h][index+1:]...)
				}
			}
		} else {
			label := inputItem[:len(inputItem)-2]
			focalLength := int(inputItem[len(inputItem)-1]) - int('0')
			h := hash(label)
			// fmt.Println("Adding", label, "to bucket", h, "with value", focalLength)
			hashMap = updateHashMap(hashMap, h, label, focalLength)
		}
	}

	total := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < len(hashMap[i]); j++ {
			total += (i + 1) * (j + 1) * hashMap[i][j].focalLength
		}
	}

	return strconv.Itoa(total)
}

func updateHashMap(hashMap [256][]HashMapItem, h int, label string, focalLength int) [256][]HashMapItem {
	for i, item := range hashMap[h] {
		if item.label == label {
			hashMap[h][i].focalLength = focalLength
			return hashMap
		}
	}
	hashMap[h] = append(hashMap[h], HashMapItem{label, focalLength})
	return hashMap
}
