package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", partOne())
	//fmt.Println("Part 2:", partTwo())
}

func mapValue(valueToBeMapped int, mappingList [][]int) int {
	for _, mapping := range mappingList {
		if valueToBeMapped > mapping[1] && valueToBeMapped < mapping[1] + mapping[2] {
			// fmt.Println(valueToBeMapped, mapping)
			return mapping[0] + (valueToBeMapped - mapping[1])
		}
	}
	return valueToBeMapped

}

func partOne() string {
	lines := readLines("input.txt")
	numberExpression := regexp.MustCompile(`(\d+)`)

	var seeds [] int
	mappings := [][][] int{}

	mappingCount := -1
	lowest := 9999999999999999

	for lineNum, line := range lines {
		if lineNum == 0 {
			seeds, _ = sliceAtoi(numberExpression.FindAllString(line, -1))
			continue
		}

		// Empty line means the mapping group changes
		if len(strings.TrimSpace(line)) == 0 {
			mappingCount += 1
			mappings = append(mappings, [][]int{})
			continue
		}

		values, _ := sliceAtoi(numberExpression.FindAllString(line, -1))
		// If no numbers it means it's the mapping group title line so we need to skip it
		if len(values) > 0 {
			mappings[mappingCount] = append(mappings[mappingCount], values)
		}
	}

	for _, seed := range seeds {
		valueToBeMapped := seed
		for _, mappingType := range mappings {
			valueToBeMapped = mapValue(valueToBeMapped, mappingType)
		}
		//fmt.Println(seed, "maps to", valueToBeMapped, "- current lowest is", lowest)
		if valueToBeMapped < lowest {
			lowest = valueToBeMapped
		}
	}

	return strconv.Itoa(lowest)
}

func partTwo() string {
	lines := readLines("input.txt")
	numberExpression := regexp.MustCompile(`(\d+)`)

	var seeds [] int
	mappings := [][][] int{}

	mappingCount := -1
	lowest := 9999999999999999

	for lineNum, line := range lines {
		if lineNum == 0 {
			seeds, _ = sliceAtoi(numberExpression.FindAllString(line, -1))
			continue
		}

		// Empty line means the mapping group changes
		if len(strings.TrimSpace(line)) == 0 {
			mappingCount += 1
			mappings = append(mappings, [][]int{})
			continue
		}

		values, _ := sliceAtoi(numberExpression.FindAllString(line, -1))
		// If no numbers it means it's the mapping group title line so we need to skip it
		if len(values) > 0 {
			mappings[mappingCount] = append(mappings[mappingCount], values)
		}
	}

	for i := 0; i < len(seeds); i+= 2 {
		for seed := seeds[i]; seed < seeds[i] + seeds[i+1]; seed++ {
			valueToBeMapped := seed
			for _, mappingType := range mappings {
				valueToBeMapped = mapValue(valueToBeMapped, mappingType)
			}
			if valueToBeMapped < lowest {
				fmt.Println(i, "--", seed, "maps to", valueToBeMapped, "- current lowest is", lowest)
				lowest = valueToBeMapped
			}
		}
	}

	return strconv.Itoa(lowest)
}
