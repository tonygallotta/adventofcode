package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

type node struct {
	id       point
	value    int
	adjacent []*node
}

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var grid, nodes, start, finish = buildGraph(lines)
	var distance = computeShortestPaths(nodes, start, len(grid[0]))
	fmt.Println("Part 1: ", distance[index(finish, len(grid[0]))])
}

func part2(lines []string) {
	var grid, nodes, _, finish = buildGraph(lines)
	var minDistance = math.MaxInt
	var checks = 0
	for _, node := range nodes {
		if node.value == 0 {
			checks += 1
			fmt.Println("check ", checks, " min distance = ", minDistance)
			var distance = computeShortestPaths(nodes, node.id, len(grid[0]))[index(finish, len(grid[0]))]
			if distance < minDistance {
				minDistance = distance
			}
		}
	}
	fmt.Println("Part 2: ", minDistance)
}

func buildGraph(lines []string) ([][]int, map[point]*node, point, point) {
	var grid = [][]int{}
	var start, finish point
	var nodes = make(map[point]*node)
	for i, line := range lines {
		var row []int
		for j, c := range line {
			var p = point{i, j}
			var value int
			if c == 'S' {
				start = p
				value = 0
			} else if c == 'E' {
				finish = p
				value = 25
			} else {
				value = int(c) - int('a')
			}
			row = append(row, value)
			nodes[p] = &node{p, value, []*node{}}
		}
		grid = append(grid, row)
	}
	for p, node := range nodes {
		var neighborPoints = neighbors(p, grid)
		for _, neighborPoint := range neighborPoints {
			var neighborNode, _ = nodes[neighborPoint]
			if neighborNode.value <= node.value+1 {
				node.adjacent = append(node.adjacent, neighborNode)
			}
		}
	}
	return grid, nodes, start, finish
}

func neighbors(p point, grid [][]int) []point {
	var gridSizeX, gridSizeY = len(grid), len(grid[0])
	var neighbors = []point{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0 || j == 0) && i != j {
				var possibleNeighbor = point{p.x + i, j + p.y}
				if possibleNeighbor.x >= 0 && possibleNeighbor.y >= 0 && possibleNeighbor.x < gridSizeX && possibleNeighbor.y < gridSizeY {
					neighbors = append(neighbors, possibleNeighbor)
				}
			}
		}
	}
	return neighbors
}

func computeShortestPaths(nodes map[point]*node, start point, gridSize int) []int {
	var unvisited = make(map[point]bool)
	var distance = make([]int, len(nodes))

	for point, _ := range nodes {
		unvisited[point] = true
		distance[index(point, gridSize)] = math.MaxInt
	}
	distance[index(start, gridSize)] = 0
	for len(unvisited) > 0 {
		var pointToVisit, pointDistance = minDistanceNode(unvisited, distance, gridSize)
		var u, _ = nodes[pointToVisit]
		delete(unvisited, u.id)
		if pointDistance == math.MaxInt {
			continue
		}
		for _, neighbor := range u.adjacent {
			if val, _ := unvisited[neighbor.id]; val {
				var newDistance = distance[index(u.id, gridSize)] + 1
				if newDistance < distance[index(neighbor.id, gridSize)] {
					distance[index(neighbor.id, gridSize)] = newDistance
				}
			}
		}
	}
	return distance
}

func minDistanceNode(unvisited map[point]bool, distance []int, gridSize int) (point, int) {
	var minDistance = math.MaxInt
	var closest point
	for p, _ := range unvisited {
		var currentDistance = distance[index(p, gridSize)]
		if currentDistance <= minDistance {
			minDistance = currentDistance
			closest = p
		}
	}
	return closest, minDistance
}

func index(p point, gridSize int) int {
	return p.x*gridSize + p.y
}
