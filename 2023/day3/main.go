package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("day3/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

type Grid struct {
	lines   []string
	numRows int
	numCols int
}

func MakeGrid(lines []string) Grid {
	return Grid{lines, len(lines), len(lines[0])}
}

func (g Grid) HasAdjacentSpecialChar(row int, colStart int, colEnd int) bool {
	var specialChars = regexp.MustCompile("[^\\d\\.]")

	if colStart > 0 {
		colStart--
		if specialChars.MatchString(g.lines[row][colStart : colStart+1]) {
			return true
		}
	}
	if colEnd < g.numCols {
		if specialChars.MatchString(g.lines[row][colEnd : colEnd+1]) {
			return true
		}
		colEnd++
	}
	if row > 0 {
		rowAbove := g.lines[row-1]
		if specialChars.MatchString(rowAbove[colStart:colEnd]) {
			return true
		}
	}
	if row < g.numRows-1 {
		rowBelow := g.lines[row+1]
		if specialChars.MatchString(rowBelow[colStart:colEnd]) {
			return true
		}
	}
	return false
}

func (g *Grid) GearRatio(row int, col int) int {
	numberPattern := regexp.MustCompile("\\d+")
	adjacentPartNumbers := []int{}
	currentRow := g.lines[row]
	if col > 0 && unicode.IsNumber(rune(currentRow[col-1])) {
		matches := numberPattern.FindAllString(g.lines[row][0:col], -1)
		if len(matches) > 0 {
			partOnLeft, _ := strconv.Atoi(matches[len(matches)-1])
			adjacentPartNumbers = append(adjacentPartNumbers, partOnLeft)
		}
	}
	if col < g.numCols && unicode.IsNumber(rune(currentRow[col+1])) {
		matches := numberPattern.FindAllString(g.lines[row][col:g.numCols], -1)
		if len(matches) > 0 {
			partOnRight, _ := strconv.Atoi(matches[0])
			adjacentPartNumbers = append(adjacentPartNumbers, partOnRight)
		}
	}
	if row > 0 {
		rowAbove := g.lines[row-1]
		matches := numberPattern.FindAllStringIndex(rowAbove, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			for i := col - 1; i <= col+1; i++ {
				if i >= startIndex && i < endIndex {
					number, _ := strconv.Atoi(rowAbove[startIndex:endIndex])
					adjacentPartNumbers = append(adjacentPartNumbers, number)
					break
				}
			}
		}
	}
	if row < g.numRows-1 {
		rowBelow := g.lines[row+1]
		matches := numberPattern.FindAllStringIndex(rowBelow, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			for i := col - 1; i <= col+1; i++ {
				if i >= startIndex && i < endIndex {
					number, _ := strconv.Atoi(rowBelow[startIndex:endIndex])
					adjacentPartNumbers = append(adjacentPartNumbers, number)
					break
				}
			}
		}
	}

	if len(adjacentPartNumbers) != 2 {
		return 0
	}
	return adjacentPartNumbers[0] * adjacentPartNumbers[1]
}

func part1(lines []string) {
	var answer = 0
	var grid = MakeGrid(lines)
	for rowNum, line := range lines {
		numberPattern := regexp.MustCompile("\\d+")
		matches := numberPattern.FindAllStringIndex(line, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			if grid.HasAdjacentSpecialChar(rowNum, startIndex, endIndex) {
				number, _ := strconv.Atoi(line[startIndex:endIndex])
				answer += number
			}
		}
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	var grid = MakeGrid(lines)
	for rowNum, line := range lines {
		gearPattern := regexp.MustCompile("\\*")
		matches := gearPattern.FindAllStringIndex(line, -1)
		for _, match := range matches {
			startIndex := match[0]
			answer += grid.GearRatio(rowNum, startIndex)
		}
	}
	fmt.Println("Part 2: ", answer)
}
