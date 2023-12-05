package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day4/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	// 42411 too high
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var answer = 0
	numberPattern := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		var cardPoints = 0
		var winningNumbers = map[string]bool{}
		_, numbers := SplitToPair(line, ": ")
		myNumbersStr, winningNumbersStr := SplitToPair(numbers, " | ")
		for _, winningNumber := range numberPattern.FindAllString(winningNumbersStr, -1) {
			winningNumbers[winningNumber] = true
		}
		for _, myNumber := range numberPattern.FindAllString(myNumbersStr, -1) {
			if winningNumbers[myNumber] {
				cardPoints = max(cardPoints*2, 1)
			}
		}
		answer += cardPoints
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	numberPattern := regexp.MustCompile(`\d+`)
	var winsByCard = map[int]int{}
	for i, line := range lines {
		var cardNum = i + 1
		var winningNumbers = map[string]bool{}
		_, numbers := SplitToPair(line, ": ")
		myNumbersStr, winningNumbersStr := SplitToPair(numbers, " | ")
		for _, winningNumber := range numberPattern.FindAllString(winningNumbersStr, -1) {
			winningNumbers[winningNumber] = true
		}
		for _, myNumber := range numberPattern.FindAllString(myNumbersStr, -1) {
			if winningNumbers[myNumber] {
				winsByCard[cardNum] = winsByCard[cardNum] + 1
			}
		}
	}

	var cardNumberCounts = map[int]int{}
	for i, _ := range lines {
		cardNumberCounts[i+1] += 1 // Count the original
		ProcessCard(i+1, winsByCard, cardNumberCounts)
	}
	for _, val := range cardNumberCounts {
		answer += val
	}
	fmt.Println("Part 2: ", answer)
}

func ProcessCard(cardNumber int, winsByCard map[int]int, cardNumberCounts map[int]int) {
	for i := cardNumber + 1; i <= cardNumber+winsByCard[cardNumber]; i++ {
		cardNumberCounts[i] += 1 // Count each copy
		ProcessCard(i, winsByCard, cardNumberCounts)
	}
}

func SplitToPair(s string, separator string) (string, string) {
	pair := strings.Split(s, separator)
	return pair[0], pair[1]
}
