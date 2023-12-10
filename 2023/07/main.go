package main

import (
	"fmt"
	"sort"
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

type Hand struct {
	hand         string
	cardCounts   [15]int
	sortableHand [5]int
	handType     int
	bet          int
	sortValue    int
}

func getFaceValue(card rune) int {
	numVal := card - '0'
	if numVal > 0 && numVal <= 9 {
		return int(numVal)
	}
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	default:
		return 10
	}
}

func parseToSortable(hand string) [5]int {
	rv := [5]int{}
	for i, card := range hand {
		rv[i] = getFaceValue(card)
	}
	return rv
}

/*
	Hand type mapping:
	1 - high card
	2 - one pair
	3 - two pairs
	4 - three of a kind
	5 - full house
	6 - four of a kind
	7 - five of a kind
*/

func getHandType(cardCounts [15]int) int {
	max := 0
	pairsOrBetter := 0
	for _, cardNum := range cardCounts {
		if cardNum > max {
			max = cardNum
		}
		if cardNum > 1 {
			pairsOrBetter++
		}
	}
	switch max {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if pairsOrBetter > 1 {
			return 5
		} else {
			return 4
		}
	case 2:
		if pairsOrBetter > 1 {
			return 3
		} else {
			return 2
		}
	default:
		return 1
	}
}

func parseHand(line string) Hand {
	parts := strings.Split(line, " ")
	sortableHand := parseToSortable(parts[0])
	cardCounts := [15]int{}
	for _, card := range sortableHand {
		cardCounts[card] += 1
	}
	handType := getHandType(cardCounts)
	bet, _ := strconv.Atoi(parts[1])
	sortValue := handType << (4 * 5)
	for i, val := range sortableHand {
		sortValue |= val << (4 * (4 - i))
	}
	//fmt.Println(line, sortableHand, cardCounts, handType, bet)
	return Hand{parts[0], cardCounts, sortableHand, handType, bet, sortValue}
}

func partOne(filename string) string {
	lines := readLines(filename)

	hands := make([]Hand, 0, len(lines))

	for _, line := range lines {
		hands = append(hands, parseHand(line))
	}

	/*
		for _, hand := range hands {
			fmt.Println(hand)
		}

		fmt.Println("Sorting...")
	*/

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].sortValue < hands[j].sortValue
	})

	total := 0
	for i, hand := range hands {
		//fmt.Println(hand)
		total += hand.bet * (i + 1)
	}

	return strconv.Itoa(total)
}
