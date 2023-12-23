package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Vertex struct {
	id       Point
	value    rune
	adjacent []*Vertex
}

func (p *Point) Index(gridSize int) int {
	return p.x*gridSize + p.y
}

func (v *Vertex) CanConnect(label string, neighbor *Vertex) bool {
	var westConnectors = map[rune]bool{'-': true, 'J': true, '7': true, 'S': true}
	var eastConnectors = map[rune]bool{'-': true, 'L': true, 'F': true, 'S': true}
	var northConnectors = map[rune]bool{'|': true, 'L': true, 'J': true, 'S': true}
	var southConnectors = map[rune]bool{'|': true, '7': true, 'F': true, 'S': true}
	//fmt.Printf("Checking %v %v of %v\n", neighbor.id, label, v.id)
	switch label {
	case "left":
		return westConnectors[v.value] && eastConnectors[neighbor.value]
	case "right":
		return eastConnectors[v.value] && westConnectors[neighbor.value]
	case "up":
		return northConnectors[v.value] && southConnectors[neighbor.value]
	case "down":
		return southConnectors[v.value] && northConnectors[neighbor.value]
	}
	return false
}

func main() {
	data, _ := os.ReadFile("day10/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var debug = false
	var answer = 0
	var grid, nodes, start = BuildGraph(lines)
	var gridSize = len(grid[0])
	var distance = ComputeShortestPaths(nodes, start, gridSize)

	if debug {
		fmt.Println("Grid size is", gridSize)
		for i := 0; i < len(grid[0]); i++ {
			for j := 0; j < len(grid); j++ {
				current, ok := nodes[Point{i, j}]
				if ok && distance[current.id.Index(gridSize)] != math.MaxInt {
					fmt.Print(string(current.value))
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			current, ok := nodes[Point{i, j}]
			if ok && distance[current.id.Index(gridSize)] != math.MaxInt {
				currentDist := distance[current.id.Index(gridSize)]
				if debug {
					fmt.Print(currentDist)
				}
				if currentDist > answer {
					answer = currentDist
				}
			} else if debug {
				fmt.Print(".")
			}
		}
		if debug {
			fmt.Println()
		}
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var debug = false
	var answer = 0
	var grid, nodes, start = BuildGraph(lines)
	var gridSize = len(grid[0])
	var distance = ComputeShortestPaths(nodes, start, len(grid[0]))

	var verticalConnectors = map[rune]bool{'|': true, 'L': true, 'J': true, '7': true, 'F': true}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			current, ok := nodes[Point{i, j}]
			isLoopNode := ok && distance[current.id.Index(gridSize)] != math.MaxInt
			if !isLoopNode {
				var intersections = 0
				var lastVertical = '|'
				// Ray casting algorithm
				for k := j - 1; k >= 0; k-- {
					var toCheck = grid[i][k]
					point := Point{i, k}
					isOnPath := distance[point.Index(gridSize)] != math.MaxInt
					if !isOnPath {
						continue
					}
					switch toCheck {
					case '|':
						intersections += 1
					case 'F':
						if lastVertical == 'J' {
							intersections += 1
						}
					case 'L':
						if lastVertical == '7' {
							intersections += 1
						}
					}
					if verticalConnectors[toCheck] {
						lastVertical = toCheck
					}
				}
				inLoop := intersections%2 == 1
				if inLoop {
					answer += 1
				}
				if debug {
					if inLoop {
						fmt.Print("I")
					} else {
						fmt.Print("O")
					}
				}
			} else if debug {
				if isLoopNode {
					fmt.Print(string(current.value))
				} else {
					fmt.Print(".")
				}
			}
		}
		if debug {
			fmt.Println()
		}
	}
	fmt.Println("Part 2: ", answer)
}

func BuildGraph(lines []string) ([][]rune, map[Point]*Vertex, Point) {
	var grid = [][]rune{}
	var start Point
	var nodes = make(map[Point]*Vertex)
	for i, line := range lines {
		var row []rune
		for j, c := range line {
			var p = Point{i, j}
			if c == 'S' {
				start = p
			}
			row = append(row, c)
			if c != '.' {
				nodes[p] = &Vertex{p, c, []*Vertex{}}
			}
		}
		grid = append(grid, row)
	}
	for p, node := range nodes {
		var neighborPoints = Neighbors(p, grid)
		for label, neighborPoint := range neighborPoints {
			var neighborNode, ok = nodes[neighborPoint]
			if ok && node.CanConnect(label, neighborNode) {
				//fmt.Printf("%v (%v) has connection %v (%v) to %v\n", node.id, string(node.value), label, string(neighborNode.value), neighborNode.id)
				node.adjacent = append(node.adjacent, neighborNode)
			}
		}
	}
	return grid, nodes, start
}

// non-diagonal only
func Neighbors(p Point, grid [][]rune) map[string]Point {
	var gridSizeX, gridSizeY = len(grid), len(grid[0])
	var neighbors = map[string]Point{}
	var labels = []string{"up", "left", "right", "down"}
	var idx = 0
	for i := -1; i <= 1; i++ {
		// (-1, 0) (0, -1) (0, 1) (1, 0)
		for j := -1; j <= 1; j++ {
			if (i == 0 || j == 0) && i != j {
				var possibleNeighbor = Point{p.x + i, j + p.y}
				if possibleNeighbor.x >= 0 && possibleNeighbor.y >= 0 && possibleNeighbor.x < gridSizeX && possibleNeighbor.y < gridSizeY {
					neighbors[labels[idx]] = possibleNeighbor
				}
				//fmt.Printf("%v,%v is %v\n", i, j, labels[idx])
				idx += 1
			}
		}
	}
	return neighbors
}

func ComputeShortestPaths(nodes map[Point]*Vertex, start Point, gridSize int) []int {
	var unvisited = make(map[Point]bool)
	var distance = make([]int, gridSize*gridSize)

	for point, _ := range nodes {
		unvisited[point] = true
		distance[point.Index(gridSize)] = math.MaxInt
	}
	distance[start.Index(gridSize)] = 0
	for len(unvisited) > 0 {
		var pointToVisit, pointDistance = minDistanceNode(unvisited, distance, gridSize)
		var u, _ = nodes[pointToVisit]
		delete(unvisited, u.id)
		if pointDistance == math.MaxInt {
			continue
		}
		for _, neighbor := range u.adjacent {
			if val, _ := unvisited[neighbor.id]; val {
				var newDistance = distance[u.id.Index(gridSize)] + 1
				if newDistance < distance[neighbor.id.Index(gridSize)] {
					distance[neighbor.id.Index(gridSize)] = newDistance
				}
			}
		}
	}
	return distance
}

func minDistanceNode(unvisited map[Point]bool, distance []int, gridSize int) (Point, int) {
	var minDistance = math.MaxInt
	var closest Point
	for p, _ := range unvisited {
		var currentDistance = distance[p.Index(gridSize)]
		if currentDistance <= minDistance {
			minDistance = currentDistance
			closest = p
		}
	}
	return closest, minDistance
}
