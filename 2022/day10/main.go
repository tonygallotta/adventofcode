package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var signalStrength = 0
	var nextObservation = 20
	var cycleAfterExecution = 0
	var registerXValue = 1
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var instr = parts[0]
		var operation = func() {}
		switch instr {
		case "addx":
			cycleAfterExecution += 2
			operation = func() {
				var value, _ = strconv.Atoi(parts[1])
				registerXValue += value
			}
		case "noop":
			cycleAfterExecution += 1
		}
		if cycleAfterExecution >= nextObservation {
			signalStrength += nextObservation * registerXValue
			nextObservation += 40
		}
		operation()
	}
	fmt.Println("Part 1: ", signalStrength)
}

// RLEZFLGE
func part2(lines []string) {
	var cycle = 1
	var registerXValue = 1
	var crt = make([]string, 240)
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var instr = parts[0]
		var operation func()
		var position = (cycle - 1) % 40
		var nextPosition = (position + 1) % 40
		if instr == "addx" {
			if registerXValue >= position-1 && registerXValue <= position+1 {
				crt[cycle-1] = "#"
			} else {
				crt[cycle-1] = "."
			}
			if registerXValue >= nextPosition-1 && registerXValue <= (nextPosition+1)%40 {
				crt[cycle] = "#"
			} else {
				crt[cycle] = "."
			}
			operation = func() {
				var value, _ = strconv.Atoi(parts[1])
				registerXValue += value
				cycle += 2
			}
		} else {
			if registerXValue >= position-1 && registerXValue <= position+1 {
				crt[cycle-1] = "#"
			} else {
				crt[cycle-1] = "."
			}
			operation = func() { cycle += 1 }
		}
		operation()
	}
	printScreen(crt)
}

func printScreen(crt []string) {
	for i, px := range crt {
		fmt.Print(px)
		if i > 0 && (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
