package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var result = 0
	for i := 0; i < len(lines)-1; i += 3 {
		var correctlyOrdered = compare(lines[i], lines[i+1])
		if correctlyOrdered < 0 {
			result += (i + 3) / 3
		}
	}
	fmt.Println("Part 1: ", result)
}

func part2(lines []string) {
	var divider1, divider2 = "[[2]]", "[[6]]"
	var linesToSort = []string{divider1, divider2}
	for _, line := range lines {
		if len(line) > 0 {
			linesToSort = append(linesToSort, line)
		}
	}
	sort.Slice(linesToSort, func(i, j int) bool {
		return compare(linesToSort[i], linesToSort[j]) < 0
	})
	var answer = 1
	for i, line := range linesToSort {
		if line == divider1 || line == divider2 {
			answer *= i + 1
		}
	}
	fmt.Println("Part 2: ", answer)
}

func compare(first, second string) int {
	var p1, p2 interface{}
	json.Unmarshal([]byte(first), &p1)
	json.Unmarshal([]byte(second), &p2)
	var p1Number, isPacket1Number = p1.(float64)
	var p2Number, isPacket2Number = p2.(float64)
	// fmt.Printf("Comparing %s %s, %T, %T\n", first, second, p1, p2)
	if isPacket1Number && isPacket2Number {
		return int(p1Number) - int(p2Number)
	} else if isPacket1Number {
		return compare(fmt.Sprintf("[%d]", int(p1Number)), second)
	} else if isPacket2Number {
		return compare(first, fmt.Sprintf("[%d]", int(p2Number)))
	} else {
		// Both lists
		var p1ListValue, _ = p1.([]interface{})
		var p2ListValue, _ = p2.([]interface{})
		if len(p1ListValue) == 0 && len(p2ListValue) > 0 {
			return -1
		} else if len(p1ListValue) > 0 && len(p2ListValue) == 0 {
			return 1
		} else if len(p1ListValue) == 0 && len(p2ListValue) == 0 {
			return 0
		}
		for i := 0; i < len(p1ListValue); i++ {
			if i >= len(p2ListValue) {
				return 1
			}
			var p1Next, _ = json.Marshal(p1ListValue[i])
			var p2Next, _ = json.Marshal(p2ListValue[i])
			var comparisonResult = compare(string(p1Next), string(p2Next))
			if comparisonResult != 0 {
				return comparisonResult
			}
		}
		var p1Next, _ = json.Marshal(p1ListValue[1:])
		var p2Next, _ = json.Marshal(p2ListValue[1:])
		return compare(string(p1Next), string(p2Next))
	}
}
