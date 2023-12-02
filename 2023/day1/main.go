package main

import (
	"fmt"
	"os"
	"strconv"
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
	var answer = 0
	for _, line := range lines {
		var lineVal = ""
		for _, char := range line {
			var _, err = strconv.Atoi(string(char))
			if err == nil {
				lineVal = lineVal + string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			var char = line[i]
			var _, err = strconv.Atoi(string(char))
			if err == nil {
				lineVal = lineVal + string(char)
				break
			}
		}
		fmt.Printf("Line val %v", lineVal)
		var calibrationValue, _ = strconv.Atoi(lineVal)
		answer = answer + calibrationValue
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	for _, line := range lines {
		var lineVal = ""
		for i := 0; i < len(line); i++ {
			var char = line[i]
			var _, err = strconv.Atoi(string(char))
			if err == nil {
				lineVal = lineVal + string(char)
				break
			} else {
				var maybeTextNumber = ExtractNumber(line[i:])
				if maybeTextNumber != "" {
					lineVal = lineVal + maybeTextNumber
					break
				}
			}
		}
		var lastNumber = ""
		for i := 0; i < len(line); i++ {
			var char = line[i]
			var _, err = strconv.Atoi(string(char))
			if err == nil {
				lastNumber = string(char)
			} else {
				var maybeTextNumber = ExtractNumber(line[i:])
				if maybeTextNumber != "" {
					lastNumber = maybeTextNumber
				}
			}
		}
		lineVal = lineVal + lastNumber
		var calibrationValue, _ = strconv.Atoi(lineVal)
		answer = answer + calibrationValue
	}
	fmt.Println("Part 2: ", answer)
}

func ExtractNumber(s string) string {
	if strings.HasPrefix(s, "one") {
		return "1"
	}
	if strings.HasPrefix(s, "two") {
		return "2"
	}
	if strings.HasPrefix(s, "three") {
		return "3"
	}
	if strings.HasPrefix(s, "four") {
		return "4"
	}
	if strings.HasPrefix(s, "five") {
		return "5"
	}
	if strings.HasPrefix(s, "six") {
		return "6"
	}
	if strings.HasPrefix(s, "seven") {
		return "7"
	}
	if strings.HasPrefix(s, "eight") {
		return "8"
	}
	if strings.HasPrefix(s, "nine") {
		return "9"
	}
	return ""
}
