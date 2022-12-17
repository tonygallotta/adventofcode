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

type path struct {
	points []point
}

type cave struct {
	bounds bounds
	paths  []path
	grid   [][]string
}

type bounds struct {
	xMin int
	yMin int
	xMax int
	yMax int
}

func (self *cave) print() {
	fmt.Println("Bounds", self.bounds)
	for i := range self.grid {
		for j := 0; j < len(self.grid[i]); j++ {
			fmt.Print(self.grid[i][j])
		}
		fmt.Println()
	}
}

func (self *cave) get(p point) string {
	return self.grid[p.y][p.x-self.bounds.xMin]
}

func (self *cave) set(p point, value string) {
	self.grid[p.y][p.x-self.bounds.xMin] = value
}

func MakeCave(lines []string) cave {
	var bounds = bounds{500, 0, 500, 0}
	var paths = []path{}
	for _, line := range lines {
		var pointStrs = strings.Split(line, " -> ")
		var points = []point{}
		for _, p := range pointStrs {
			var coords = strings.Split(p, ",")
			var x, _ = strconv.Atoi(coords[0])
			var y, _ = strconv.Atoi(coords[1])
			points = append(points, point{x, y})
			if x < bounds.xMin {
				bounds.xMin = x
			}
			if x > bounds.xMax {
				bounds.xMax = x
			}
			if y < bounds.yMin {
				bounds.yMin = y
			}
			if y > bounds.yMax {
				bounds.yMax = y
			}
		}
		paths = append(paths, path{points})
	}

	var grid = make([][]string, bounds.yMax-bounds.yMin+1)
	for i := range grid {
		var length = bounds.xMax - bounds.xMin + 1
		grid[i] = make([]string, length)
		for j := 0; j < length; j++ {
			grid[i][j] = AIR
		}
	}
	for _, path := range paths {
		// look at each path's points pairwise to fill out the grid
		for i := 0; i < len(path.points)-1; i++ {
			var start, end = path.points[i], path.points[i+1]
			if start.x == end.x {
				for y := min(start.y, end.y); y <= max(start.y, end.y); y++ {
					grid[y][start.x-bounds.xMin] = ROCK
				}
			}
			if start.y == end.y {
				for x := min(start.x, end.x); x <= max(start.x, end.x); x++ {
					grid[start.y][x-bounds.xMin] = ROCK
				}
			}
		}
	}
	return cave{bounds, paths, grid}
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

const AIR = "."
const ROCK = "#"
const SAND = "o"

func part1(lines []string) {
	var result = 0
	var cave = MakeCave(lines)
	cave.print()
out:
	for {
		var sandPoint = point{500, 0}
		cave.set(sandPoint, SAND)
		var moved = true
		for moved {
			var possibleNext = []point{{sandPoint.x, sandPoint.y + 1}, {sandPoint.x - 1, sandPoint.y + 1}, {sandPoint.x + 1, sandPoint.y + 1}}
			moved = false
			for _, point := range possibleNext {
				if point.y > cave.bounds.yMax || point.x > cave.bounds.xMax || point.x < cave.bounds.xMin {
					break out
				} else if cave.get(point) == AIR {
					cave.set(point, SAND)
					cave.set(sandPoint, AIR)
					sandPoint = point
					moved = true
					break
				}
			}
		}
		// Came to rest
		result += 1
	}
	cave.print()
	fmt.Println("Part 1: ", result)
}

func part2(lines []string) {
	var result = 0
	var initialCave = MakeCave(lines)
	// This is hacky, I increased the expansion amount until I got a stable answer. For the sample input, an expansion of 2 was required.
	var expansion = 220
	var cave = MakeCave(append(lines, fmt.Sprintf("%d,%d -> %d,%d", initialCave.bounds.xMin-expansion, initialCave.bounds.yMax+2, initialCave.bounds.xMax+expansion, initialCave.bounds.yMax+2)))
	cave.print()
	var leftOverflow, rightOverflow = 0, 0
	var leftOverflowLevel, rightOverflowLevel = 0, 0
out:
	for {
		var sandPoint = point{500, 0}
		cave.set(sandPoint, SAND)
		var moved = true
		for moved {
			var possibleNext = []point{{sandPoint.x, sandPoint.y + 1}, {sandPoint.x - 1, sandPoint.y + 1}, {sandPoint.x + 1, sandPoint.y + 1}}
			moved = false
			for _, point := range possibleNext {
				if point.y > cave.bounds.yMax || point.x > cave.bounds.xMax || point.x < cave.bounds.xMin {
					break
				} else if cave.get(point) == AIR {
					cave.set(point, SAND)
					cave.set(sandPoint, AIR)
					sandPoint = point
					moved = true
					break
				}
			}
			// Try left and right once more allowing for overflow
			if sandPoint.x == cave.bounds.xMin {
				// We can overflow until we fill up the next triangle on the left
				if leftOverflow < (leftOverflowLevel+1)/2 {
					// fmt.Println("Overflowing on left @ ", point)
					leftOverflow += 1
					result += 1
					cave.set(sandPoint, AIR)
					moved = true
					break
				}
				if leftOverflow == (leftOverflowLevel+1)/2 {
					leftOverflowLevel += 1
				}
			} else if sandPoint.x == cave.bounds.xMax {
				// We can overflow until we fill up the next triangle on the right
				if rightOverflow < (rightOverflowLevel+1)/2 {
					// fmt.Println("Overflowing on right @ level ", level)
					rightOverflow += 1
					result += 1
					cave.set(sandPoint, AIR)
					moved = true
					break
				}
				if rightOverflow == (rightOverflowLevel+1)/2 {
					rightOverflowLevel += 1
				}
			}
		}
		// Came to rest
		result += 1
		if sandPoint.y == 0 {
			// Didn't move from the origin
			break out
		}
	}
	cave.print()
	fmt.Println("Part 2:", result)
}
