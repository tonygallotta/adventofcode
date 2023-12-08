package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day8/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

type Node struct {
	label string
	left  *Node
	right *Node
}

func part1(lines []string) {
	var nodesByLabel = map[string]*Node{}
	var directions = []rune(lines[0])
	for _, line := range lines[2:] {
		label, _ := SplitToPair(line, " = ")
		nodesByLabel[label] = &Node{label, nil, nil}
	}
	for _, line := range lines[2:] {
		label, adjacent := SplitToPair(line, " = ")
		leftLabel, rightLabel := SplitToPair(strings.ReplaceAll(strings.ReplaceAll(adjacent, "(", ""), ")", ""), ", ")
		node := nodesByLabel[label]
		leftNode := nodesByLabel[leftLabel]
		rightNode := nodesByLabel[rightLabel]
		node.left = leftNode
		node.right = rightNode
	}
	var steps = 0
	for current := nodesByLabel["AAA"]; current.label != "ZZZ"; {
		move := directions[steps%len(directions)]
		switch move {
		case 'L':
			current = current.left
			break
		case 'R':
			current = current.right
			break
		}
		steps += 1
	}
	fmt.Println("Part 1: ", steps)
}

func part2(lines []string) {
	var nodesByLabel = map[string]*Node{}
	var directions = []rune(lines[0])
	for _, line := range lines[2:] {
		label, _ := SplitToPair(line, " = ")
		nodesByLabel[label] = &Node{label, nil, nil}
	}
	for _, line := range lines[2:] {
		label, adjacent := SplitToPair(line, " = ")
		leftLabel, rightLabel := SplitToPair(strings.ReplaceAll(strings.ReplaceAll(adjacent, "(", ""), ")", ""), ", ")
		node := nodesByLabel[label]
		leftNode := nodesByLabel[leftLabel]
		rightNode := nodesByLabel[rightLabel]
		node.left = leftNode
		node.right = rightNode
	}
	var firstZs = []int{}
	for _, node := range NodesEndingWith(nodesByLabel, 'A') {
		for steps := 0; ; steps += 1 {
			move := directions[steps%len(directions)]
			switch move {
			case 'L':
				node = node.left
				break
			case 'R':
				node = node.right
				break
			}
			if rune(node.label[2]) == 'Z' {
				firstZs = append(firstZs, steps+1)
				break
			}
		}
	}
	fmt.Println("Part 2: ", LeastCommonMultiple(firstZs))
}

func SplitToPair(s string, separator string) (string, string) {
	pair := strings.Split(s, separator)
	return pair[0], pair[1]
}

func NodesEndingWith(nodesByLabel map[string]*Node, c rune) map[string]*Node {
	var result = map[string]*Node{}
	for label, node := range nodesByLabel {
		//fmt.Println(label, "ends with", string(c), "?", rune(label[2]) == c)
		if rune(label[2]) == c {
			result[label] = node
		}
	}
	return result
}

func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}

func LeastCommonMultiple(arr []int) int {
	ans := arr[0]
	for i := 1; i < len(arr); i += 1 {
		ans = (arr[i] * ans) / GreatestCommonDivisor(arr[i], ans)
	}
	return ans
}
