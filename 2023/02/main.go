package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}

func partOne() string {
	lines := readLines("input.txt")
	game, _ := regexp.Compile(`^Game (\d+): (.*)$`)
	showing, _ := regexp.Compile(`(\d+) (\w+)`)
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}
	total := 0

	for _, line := range lines {
		gameParts := game.FindStringSubmatch(line)
		gameId, _ := strconv.Atoi(gameParts[1])
		gameIsOkay := true
		sets := strings.Split(gameParts[2], ";")
		for _, set := range sets {
			showingItems := showing.FindAllStringSubmatch(set, -1)
			for _, showingItem := range showingItems {
				showingItemValue, _ := strconv.Atoi(showingItem[1])
				if showingItemValue > limits[showingItem[2]] {
					gameIsOkay = false
					break
				}
			}
			if !gameIsOkay {
				break
			}
		}
		if gameIsOkay {
			total += gameId
		}
	}
	return strconv.Itoa(total)
}

func partTwo() string {
	lines := readLines("input.txt")
	game, _ := regexp.Compile(`^Game (\d+): (.*)$`)
	showing, _ := regexp.Compile(`(\d+) (\w+)`)
	total := 0

	for _, line := range lines {
		limits := map[string]int{"red": 0, "green": 0, "blue": 0}
		gameParts := game.FindStringSubmatch(line)
		//		gameId, _ := strconv.Atoi(gameParts[1])
		sets := strings.Split(gameParts[2], ";")
		for _, set := range sets {
			showingItems := showing.FindAllStringSubmatch(set, -1)
			for _, showingItem := range showingItems {
				showingItemValue, _ := strconv.Atoi(showingItem[1])
				if showingItemValue > limits[showingItem[2]] {
					limits[showingItem[2]] = showingItemValue
				}
			}
		}
		total += limits["red"] * limits["blue"] * limits["green"]
	}
	return strconv.Itoa(total)
}
