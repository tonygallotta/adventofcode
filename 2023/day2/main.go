package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day2/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var answer = 0
	var maxesByColor = map[string]int{"red": 12, "green": 13, "blue": 14}
	numberPattern := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		gameNumberString := numberPattern.FindString(line)
		gameNumber, _ := strconv.Atoi(gameNumberString)
		setsStart := len("Game :") + len(gameNumberString)
		var valid = true
	outerLoop:
		for _, cubeSet := range strings.Split(line[setsStart:], ";") {
			for _, cubeAndCount := range strings.Split(cubeSet, ", ") {
				countAsStr, color := SplitToPair(strings.TrimSpace(cubeAndCount), " ")
				count, _ := strconv.Atoi(countAsStr)
				if count > maxesByColor[color] {
					valid = false
					break outerLoop
				}
			}
		}
		if valid {
			answer += gameNumber
		}
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	for _, line := range lines {
		_, gameResult := SplitToPair(line, ": ")
		var maxByColor = map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, cubeSet := range strings.Split(gameResult, ";") {
			for _, cubeAndCount := range strings.Split(cubeSet, ", ") {
				countAsStr, color := SplitToPair(strings.TrimSpace(cubeAndCount), " ")
				count, _ := strconv.Atoi(countAsStr)
				maxByColor[color] = max(count, maxByColor[color])
			}
		}
		answer += maxByColor["red"] * maxByColor["green"] * maxByColor["blue"]
	}
	fmt.Println("Part 2: ", answer)
}

func SplitToPair(s string, separator string) (string, string) {
	pair := strings.Split(s, separator)
	return pair[0], pair[1]
}
