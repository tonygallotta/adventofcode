package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day7/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var answer = 0
	var scores = map[string]int{}
	var bids = map[string]int{}
	for _, line := range lines {
		hand, bidAsStr := SplitToPair(line, " ")
		var score = ScoreHand(hand, false)
		scores[line] = score
		bids[line] = ParseInt(bidAsStr)
	}
	sort.Slice(lines, func(i, j int) bool {
		return scores[lines[i]] < scores[lines[j]]
	})
	for i, line := range lines {
		rank := i + 1
		//fmt.Printf("%v=%v, bid %v | score %v\n", rank, line, bids[line], scores[line])
		answer += rank * bids[line]
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	var scores = map[string]int{}
	var bids = map[string]int{}
	for _, line := range lines {
		hand, bidAsStr := SplitToPair(line, " ")
		var score = ScoreHand(hand, true)
		scores[line] = score
		bids[line] = ParseInt(bidAsStr)
	}
	sort.Slice(lines, func(i, j int) bool {
		return scores[lines[i]] < scores[lines[j]]
	})
	for i, line := range lines {
		rank := i + 1
		//fmt.Printf("%v=%v, bid %v | score %v\n", rank, line, bids[line], scores[line])
		answer += rank * bids[line]
	}
	fmt.Println("Part 2: ", answer) // 249104563 too low
}

// 7 - Five of a kind, where all five cards have the same label: AAAAA
// 6 - Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 5 - Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4- Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 3 - Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 2 - One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 1 - High card, where all cards' labels are distinct: 23456
func ScoreHand(hand string, jokers bool) int {
	replacements := map[rune]rune{'A': 'E', 'K': 'D', 'Q': 'C', 'J': 'B', 'T': 'A'}
	if jokers {
		replacements['J'] = '1'
	}
	countByCard := map[rune]int{}
	hexScore := ""

	for _, runeVal := range []rune(hand) {
		replacement, found := replacements[runeVal]
		if found {
			runeVal = replacement
		}
		countByCard[runeVal] += 1
		hexScore += string(runeVal)
	}
	var pairCount = 0
	var maxOfAKind = 1
	var prefix = '1'
	for k, v := range countByCard {
		if jokers && k == '1' {
			continue // Handle jokers separately
		}
		if v == 2 {
			pairCount += 1
		}
		if v > maxOfAKind {
			maxOfAKind = v
		}
	}
	jokerCount := countByCard['1']
	fmt.Println("Joker count is ", jokerCount)
	switch maxOfAKind {
	case 5:
		prefix = '7'
		break
	case 4:
		prefix = '6'
		if jokers && jokerCount == 1 {
			prefix = '7'
		}
		break
	case 3:
		if jokers && jokerCount > 0 {
			if jokerCount == 2 {
				prefix = '7'
			} else if jokerCount == 1 {
				prefix = '6' // We can use any card to make it a full house
			}
		} else if pairCount == 1 {
			prefix = '5'
		} else {
			prefix = '4'
		}
		break
	case 2:
		if jokers && jokerCount > 0 {
			if jokerCount == 3 {
				prefix = '7'
			} else if jokerCount == 2 {
				prefix = '6'
			} else if jokerCount == 1 {
				if pairCount == 2 {
					prefix = '5' // make a full house
				} else {
					prefix = '4' // make 3 of a kind
				}
			}
		} else if pairCount == 2 {
			prefix = '3'
		} else {
			prefix = '2'
		}
		break
	default:
		if jokers && jokerCount > 0 {
			switch jokerCount {
			case 5:
				prefix = '7'
				break
			case 4:
				prefix = '7'
				break
			case 3:
				prefix = '6' // make 4 of a kind
				break
			case 2:
				prefix = '4' // 3 of a kind
				break
			case 1:
				prefix = '2' // pair
				break
			}
		} else {
			prefix = '1'
		}
	}
	hexScore = string(prefix) + hexScore
	n := new(big.Int)
	n.SetString(hexScore, 16)
	fmt.Printf("Hex score for %v = %v\n", hand, hexScore)
	return int(n.Int64())
}

func SplitToPair(s string, separator string) (string, string) {
	pair := strings.Split(s, separator)
	return pair[0], pair[1]
}
func ParseInt(s string) int {
	intValue, _ := strconv.Atoi(strings.TrimSpace(s))
	return intValue
}
