package main

import (
	"fmt"
	"os"
	"sort"
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
	var acc = 0
	var max = 0
	for _, line := range lines {
		if line == "" {
			if acc > max {
				max = acc
			}
			acc = 0
		} else {
			intValue, _ := strconv.Atoi(line)
			acc += intValue
		}
	}
	fmt.Println("Part 1: ", max)
}

func part2(lines []string) {
	var acc = 0
	var amounts []int = []int{}
	for _, line := range lines {
		if line == "" {
			amounts = append(amounts, acc)
			acc = 0
		} else {
			intValue, _ := strconv.Atoi(line)
			acc += intValue
		}
	}
	sort.Ints(amounts)
	var numElves = len(amounts)
	fmt.Println("Part 2: ", (amounts[numElves-1] + amounts[numElves-2] + amounts[numElves-3]))
}
