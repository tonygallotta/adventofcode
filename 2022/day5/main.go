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
	var stacks = make([][]string, 9)
	for i, line := range lines {
		if i < 8 {
			var items = strings.Split(strings.ReplaceAll(line, "    ", " [x]"), " ")
			for j, item := range items {
				if item != "[x]" {
					stacks[j] = append(stacks[j], string(item[1]))
				}
			}
		} else if i == 9 {
			// We put the items in in the wrong order
			for i, stack := range stacks {
				stacks[i] = reverse(stack)
			}
		} else if i > 9 {
			var instruction = strings.Split(line, " ")
			var count, _ = strconv.Atoi(instruction[1])
			var fromStackNum, _ = strconv.Atoi(instruction[3])
			var toStackNum, _ = strconv.Atoi(instruction[5])
			var fromStack = stacks[fromStackNum-1]
			var toStack = stacks[toStackNum-1]
			stacks[toStackNum-1] = append(toStack, reverse(fromStack[len(fromStack)-count:])...)
			stacks[fromStackNum-1] = fromStack[:len(fromStack)-count]
		}
	}
	var result = ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			result = result + stack[len(stack)-1]
		}
	}
	fmt.Println("Part 1: ", result)
}

func part2(lines []string) {
	var count int
	fmt.Println("Part 2: ", count)
}

func reverse(src []string) []string {
	var dest = []string{}
	for i := len(src) - 1; i >= 0; i-- {
		dest = append(dest, src[i])
	}
	return dest
}
