package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var sum int
	for _, line := range lines {
		var mid = len(line) / 2
		var sack1 = line[0:mid]
		var sack2 = line[mid:]
		for _, item := range sack1 {
			if strings.Contains(sack2, string(item)) {
				sum += score(item)
				break
			}
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2(lines []string) {
	var sum int
	for i := 0; i < len(lines)-2; i += 3 {
		var sack1 = lines[i]
		var sack2 = lines[i+1]
		var sack3 = lines[i+2]
		for _, item := range sack1 {
			if strings.Contains(sack2, string(item)) && strings.Contains(sack3, string(item)) {
				sum += score(item)
				fmt.Println(i, item)
				break
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}

func score(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - int('A') + 27
	} else {
		return int(item) - int('a') + 1
	}
}
