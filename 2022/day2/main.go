package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var total = 0
	for _, line := range lines {
		var chars = []rune(line)
		var opponentThrow = chars[0]
		var myThrow = chars[2]
		var shapeScore = int(myThrow) - int('W')
		total += shapeScore + outcomeScore(opponentThrow, myThrow)
	}
	fmt.Println("Part 1: ", total)
}

// ROCK = 0
// PAPER = 1
// SCISSORS = 2
var outcomes = [3][3]int{
	{3, 6, 0},
	{0, 3, 6},
	{6, 0, 3},
}

func outcomeScore(opponentThrow, myThrow rune) int {
	var opponentThrowValue = int(opponentThrow) - int('A')
	var myThrowValue = int(myThrow) - int('X')
	return outcomes[opponentThrowValue][myThrowValue]
}

func part2(lines []string) {
	var total = 0
	for _, line := range lines {
		var chars = []rune(line)
		var opponentThrow = chars[0]
		var myThrow = throwForOutcome(opponentThrow, chars[2])
		var shapeScore = int(myThrow) - int('W')
		total += shapeScore + outcomeScore(opponentThrow, myThrow)
	}
	fmt.Println("Part 2: ", total)
}

func throwForOutcome(opponentThrow, desiredOutcome rune) rune {
	var opponentThrowValue = int(opponentThrow) - int('A')
	const LOSE = rune('X')
	const DRAW = rune('Y')
	const WIN = rune('Z')
	var requiredOutcomeScore int

	if desiredOutcome == LOSE {
		requiredOutcomeScore = 0
	} else if desiredOutcome == DRAW {
		requiredOutcomeScore = 3
	} else {
		requiredOutcomeScore = 6
	}
	for idx, outcomeScore := range outcomes[opponentThrowValue] {
		if outcomeScore == requiredOutcomeScore {
			return rune(int('X') + idx)
		}
	}
	panic("Should never get here")
}
