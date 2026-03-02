package camelcards

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

var cardToValuePart1 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardToValuePart2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 2,
	'T': 11,
	'9': 10,
	'8': 9,
	'7': 8,
	'6': 7,
	'5': 6,
	'4': 5,
	'3': 4,
	'2': 3,
}

type CamelCardsHand struct {
	OriginalHand string
	NumericHand  []int
	Bid          int
	HandType     int
}

func typeFromFrequencyMap(freq map[rune]int) int {
	size := len(freq)
	if size == 1 {
		return FIVE_OF_A_KIND
	}

	if size == 2 {
		for _, count := range freq {
			if count == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}

	if size == 3 {
		for _, count := range freq {
			if count == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	}

	if size == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func ComputeTypeOfHandPart1(hand string) int {
	freq := make(map[rune]int)
	for _, card := range hand {
		freq[card]++
	}
	return typeFromFrequencyMap(freq)
}

func ComputeTypeOfHandPart2(hand string) int {
	freq := make(map[rune]int)
	for _, card := range hand {
		freq[card]++
	}

	if len(freq) == 1 {
		return FIVE_OF_A_KIND
	}

	countJ, hasJ := freq['J']
	if !hasJ {
		return typeFromFrequencyMap(freq)
	}

	delete(freq, 'J')

	if len(freq) == 1 {
		return FIVE_OF_A_KIND
	}

	if len(freq) == 2 {
		maxCount := 0
		for _, count := range freq {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount+countJ == 4 {
			return FOUR_OF_A_KIND
		}
		return FULL_HOUSE
	}

	if len(freq) == 3 {
		maxCount := 0
		for _, count := range freq {
			if count > maxCount {
				maxCount = count
			}
		}
		if maxCount+countJ == 3 {
			return THREE_OF_A_KIND
		}
		return TWO_PAIR
	}

	if len(freq) == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func ParseInput(filename string, part1 bool) []CamelCardsHand {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	hands := make([]CamelCardsHand, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Fatalf("invalid input line: %s", line)
		}

		handStr := fields[0]
		if len(handStr) != 5 {
			log.Fatalf("invalid hand length: %s", handStr)
		}

		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("invalid bid %q: %v", fields[1], err)
		}

		numeric := make([]int, 0, 5)
		for _, card := range handStr {
			if part1 {
				numeric = append(numeric, cardToValuePart1[card])
			} else {
				numeric = append(numeric, cardToValuePart2[card])
			}
		}

		handType := ComputeTypeOfHandPart1(handStr)
		if !part1 {
			handType = ComputeTypeOfHandPart2(handStr)
		}

		hands = append(hands, CamelCardsHand{
			OriginalHand: handStr,
			NumericHand:  numeric,
			Bid:          bid,
			HandType:     handType,
		})
	}

	return hands
}

func CompareHands(handA CamelCardsHand, handB CamelCardsHand) int {
	if handA.HandType > handB.HandType {
		return 1
	}
	if handA.HandType < handB.HandType {
		return -1
	}

	for i := 0; i < len(handA.NumericHand); i++ {
		if handA.NumericHand[i] > handB.NumericHand[i] {
			return 1
		}
		if handA.NumericHand[i] < handB.NumericHand[i] {
			return -1
		}
	}

	return 0
}

func SortHandsAscending(hands []CamelCardsHand) []CamelCardsHand {
	sort.Slice(hands, func(i int, j int) bool {
		return CompareHands(hands[i], hands[j]) < 0
	})
	return hands
}

func ComputeTotalWinnings(hands []CamelCardsHand) int {
	total := 0
	sortedHands := SortHandsAscending(hands)
	for i, hand := range sortedHands {
		total += hand.Bid * (i + 1)
	}
	return total
}

func SolvePart1(filename string) int {
	hands := ParseInput(filename, true)
	return ComputeTotalWinnings(hands)
}

func SolvePart2(filename string) int {
	hands := ParseInput(filename, false)
	return ComputeTotalWinnings(hands)
}
