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
	var grid [][]int = readGrid(lines)
	var visible = make(map[point]int)
	// visible from left
	for i := 0; i < len(grid); i++ {
		var rowMax = -1
		for j := 0; j < len(grid[i]); j++ {
			val := grid[i][j]
			if val > rowMax {
				visible[point{i, j}] = val
				rowMax = val
			}
		}
	}
	// visible from right
	for i := 0; i < len(grid); i++ {
		var rowMax = -1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			val := grid[i][j]
			if val > rowMax {
				visible[point{i, j}] = val
				rowMax = val
			}
		}
	}

	// visible from top
	for i := 0; i < len(grid[0]); i++ {
		var colMax = -1
		for j := 0; j < len(grid); j++ {
			val := grid[j][i]
			if val > colMax {
				visible[point{j, i}] = val
				colMax = val
			}
		}
	}
	// visible from bottom
	for i := 0; i < len(grid[0]); i++ {
		var colMax = -1
		for j := len(grid) - 1; j > 0; j-- {
			val := grid[j][i]
			if val > colMax {
				visible[point{j, i}] = val
				colMax = val
			}
		}
	}
	fmt.Println("Part 1: ", len(visible))
}

func part2(lines []string) {
	var grid = readGrid(lines)
	var maxScore = 0
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			var score = scenicScore(grid, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("Part 2: ", maxScore)
}

func readGrid(lines []string) [][]int {
	var result = make([][]int, len(lines))
	for i, line := range lines {
		var values = make([]int, len(line))
		for j, char := range line {
			var value, _ = strconv.Atoi(string(char))
			values[j] = value
		}
		result[i] = values
	}
	return result
}

func scenicScore(grid [][]int, x, y int) int {
	var up, down, left, right int
	// up
	var threshold = grid[x][y]
	for i := x - 1; i >= 0; i-- {
		up++
		if grid[i][y] >= threshold {
			break
		}
	}
	// down
	for i := x + 1; i < len(grid); i++ {
		down++
		if grid[i][y] >= threshold {
			break
		}
	}
	// left
	for i := y - 1; i >= 0; i-- {
		left++
		if grid[x][i] >= threshold {
			break
		}
	}
	// right
	for i := y + 1; i < len(grid[x]); i++ {
		right++
		if grid[x][i] >= threshold {
			break
		}
	}
	return up * down * left * right
}
