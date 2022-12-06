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
	fmt.Println("Part 1: ", processSignal(lines[0], 4))
}

func part2(lines []string) {
	fmt.Println("Part 2: ", processSignal(lines[0], 14))
}

func processSignal(line string, markerLen int) int {
	var lastChars = []rune{}
	for i, char := range []rune(line) {
		lastChars = append(lastChars, char)
		if i >= markerLen-1 {
			if !hasDuplicates(lastChars) {
				return i + 1
			}
			lastChars = lastChars[1:]
		}
	}
	panic("Marker not found")
}

func hasDuplicates(chars []rune) bool {
	var uniques = make(map[rune]bool)
	for _, c := range chars {
		uniques[c] = true
	}
	return len(uniques) != len(chars)
}
