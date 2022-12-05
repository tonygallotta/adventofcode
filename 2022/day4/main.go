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
	var count int
	for _, line := range lines {
		var sections = strings.Split(line, ",")
		var first = toInts(strings.Split(sections[0], "-"))
		var second = toInts(strings.Split(sections[1], "-"))
		if fullyContains(first, second) || fullyContains(second, first) {
			count += 1
		}
	}
	fmt.Println("Part 1: ", count)
}

func part2(lines []string) {
	var count int
	for _, line := range lines {
		var sections = strings.Split(line, ",")
		var first = toInts(strings.Split(sections[0], "-"))
		var second = toInts(strings.Split(sections[1], "-"))
		if overlaps(first, second) || overlaps(second, first) {
			fmt.Println("Overlap in ", line)
			count += 1
		}
	}
	fmt.Println("Part 2: ", count)
}

func fullyContains(first, second []int) bool {
	return first[0] >= second[0] && first[1] <= second[1]
}

func overlaps(first, second []int) bool {
	return (first[0] >= second[0] && first[0] <= second[1]) || (first[1] >= second[0] && first[1] <= second[1])
}

func toInts(strs []string) []int {
	var result = []int{}
	for _, str := range strs {
		var intValue, _ = strconv.Atoi(str)
		result = append(result, intValue)
	}
	return result
}
