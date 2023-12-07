package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day6/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}
func part1(lines []string) {
	var answer = 1
	// d = (t-x)*x
	// d = (7-1)*1
	// d = (7-2)*2
	// d = tx - x^2
	// 9 < 7x-x^2
	// x^2-7x+9 < 0
	// x^2-15x+40 < 0
	// 1.697 < x < 5.303 (2, 3, 4, 5)
	//
	var numberPattern = regexp.MustCompile(`\d+`)
	var times = MapToInts(numberPattern.FindAllString(lines[0], -1))
	var distances = MapToInts(numberPattern.FindAllString(lines[1], -1))
	for i := 0; i < len(times); i++ {
		root1, root2 := QuadraticRoots(1, -times[i], distances[i])
		fmt.Println("Roots:", root1, root2)
		waysToWin := int(math.Floor(root1-.00001)-math.Ceil(root2+.00001)) + 1
		fmt.Printf("%v ways to win race %v\n", waysToWin, i)
		answer *= waysToWin
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 1
	var numberPattern = regexp.MustCompile(`\d+`)
	var times = MapToInts(numberPattern.FindAllString(strings.ReplaceAll(lines[0], " ", ""), -1))
	var distances = MapToInts(numberPattern.FindAllString(strings.ReplaceAll(lines[1], " ", ""), -1))
	for i := 0; i < len(times); i++ {
		root1, root2 := QuadraticRoots(1, -times[i], distances[i])
		fmt.Println("Roots:", root1, root2)
		waysToWin := int(math.Floor(root1-.00001)-math.Ceil(root2+.00001)) + 1
		fmt.Printf("%v ways to win race %v\n", waysToWin, i)
		answer *= waysToWin
	}
	fmt.Println("Part 2: ", answer)
}

func QuadraticRoots(a, b, c int) (float64, float64) {
	fmt.Printf("Roots of %vx^2 + %vx + %v", a, b, c)
	var rootPart = math.Sqrt(float64(b*b - 4*a*c))
	var root1 = (float64(-b) + rootPart) / float64(2*a)
	var root2 = (float64(-b) - rootPart) / float64(2*a)
	return root1, root2
}

func SplitToPair(s string, separator string) (string, string) {
	pair := strings.Split(s, separator)
	return pair[0], pair[1]
}

func MapToInts(s []string) []int {
	var result = []int{}

	for _, i := range s {
		j, _ := strconv.Atoi(i)
		result = append(result, j)
	}
	return result
}
