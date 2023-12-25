package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day9/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines, false)
	part2(lines)
}

func part1(lines []string, debug bool) {
	var answer = 0
	for _, line := range lines {
		var series = []int{}
		for _, str := range strings.Split(line, " ") {
			series = append(series, ParseInt(str))
		}
		var polynomial = ComputePolynomial(series)
		if debug {
			fmt.Printf("Given: %v\nComputed: ", line)
			for i, _ := range series {
				fmt.Printf("%v ", polynomial(i))
			}
			fmt.Print(polynomial(len(series)))
			fmt.Println()
		}
		answer += polynomial(len(series))
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = 0
	for _, line := range lines {
		var series = []int{}
		for _, str := range strings.Split(line, " ") {
			series = append(series, ParseInt(str))
		}
		var polynomial = ComputePolynomial(series)
		answer += polynomial(-1)
	}
	fmt.Println("Part 2: ", answer)
}

// Based off the finite difference method described here:
// https://mathworld.wolfram.com/FiniteDifference.html
func ComputePolynomial(series []int) func(int) int {
	var coefficients = []int{}
	var differenceAllZeroes = false
	var currentSeries = series

	for !differenceAllZeroes {
		var nextSeries = []int{}
		coefficients = append(coefficients, currentSeries[0])
		for i := 0; i < len(currentSeries)-1; i++ {
			nextSeries = append(nextSeries, currentSeries[i+1]-currentSeries[i])
		}
		differenceAllZeroes = IsAllZeroes(nextSeries)
		currentSeries = nextSeries
	}

	return func(n int) int {
		var f = 0
		for k, coeff := range coefficients {
			f += coeff * Choose(n, k)
		}
		return f
	}
}

func IsAllZeroes(a []int) bool {
	return slices.Min(a) == 0 && slices.Max(a) == 0
}

func Choose(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n * Choose(n-1, k-1)) / k
}

func ParseInt(s string) int {
	intValue, _ := strconv.Atoi(strings.TrimSpace(s))
	return intValue
}
