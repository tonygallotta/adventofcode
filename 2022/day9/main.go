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
	var visited = make(map[point]bool)
	var headPosition = point{0, 0}
	var tailPosition = point{0, 0}
	visited[tailPosition] = true
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var dir = parts[0]
		var amount, _ = strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			headPosition = move(headPosition, dir, 1)
			tailPosition = adjustTail(headPosition, tailPosition)
			visited[tailPosition] = true
		}
	}
	fmt.Println("Part 1: ", len(visited))
}

func part2(lines []string) {
	var visited = make(map[point]bool)
	var knots = make([]point, 10)
	visited[knots[9]] = true
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var dir = parts[0]
		var amount, _ = strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			knots[0] = move(knots[0], dir, 1)
			for j := 0; j < 9; j++ {
				knots[j+1] = adjustTail(knots[j], knots[j+1])
			}
			visited[knots[9]] = true
		}
	}
	fmt.Println("Part 2: ", len(visited))
}

func move(from point, dir string, amount int) point {
	switch dir {
	case "U":
		return point{from.x, from.y + amount}
	case "D":
		return point{from.x, from.y - amount}
	case "L":
		return point{from.x - amount, from.y}
	case "R":
		return point{from.x + amount, from.y}
	}
	panic("Unknown direction " + dir)
}

func adjustTail(headPosition, tailPosition point) point {
	var xDistance = headPosition.x - tailPosition.x
	var yDistance = headPosition.y - tailPosition.y
	if abs(xDistance) == 2 && yDistance == 0 {
		return point{tailPosition.x + signum(xDistance), tailPosition.y}
	} else if abs(yDistance) == 2 && xDistance == 0 {
		return point{tailPosition.x, tailPosition.y + signum(yDistance)}
	} else if tailPosition.x != headPosition.x && tailPosition.y != headPosition.y && abs(xDistance)+abs(yDistance) > 2 {
		return point{tailPosition.x + signum(xDistance), tailPosition.y + signum(yDistance)}
	}
	return tailPosition
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func signum(val int) int {
	if val < 0 {
		return -1
	}
	return 1
}
