package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []uint64
	operation func(old uint64) uint64
	// Returns the monkey number to throw to
	test func(worryLevel uint64) int
}

func (self *Monkey) inspect(itemNumber int) {

}

func main() {
	data, _ := os.ReadFile("sample_input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var monkeys = []*Monkey{}
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, &Monkey{parseItems(lines[i+1]), parseOperation(lines[i+2]), parseTest(lines[i+3 : i+6])})
	}
	var inspections = make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for monkeyNum, monkey := range monkeys {
			for _, item := range monkey.items {
				// fmt.Println("Inspecting item with worry level ", item)
				var worryLevel = monkey.operation(item) / 3
				var throwTo = monkey.test(worryLevel)
				// fmt.Println("Item with worry level ", worryLevel, " is thrown to monkey ", throwTo)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
				monkey.items = monkey.items[1:]
				inspections[monkeyNum] += 1
			}
		}
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println("Part 1: ", inspections[0]*inspections[1])
}

func part2(lines []string) {
	var monkeys = []*Monkey{}
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, &Monkey{parseItems(lines[i+1]), parseOperation(lines[i+2]), parseTest(lines[i+3 : i+6])})
	}
	var inspections = make([]int, len(monkeys))
	for i := 0; i < 10_000; i++ {
		for monkeyNum, monkey := range monkeys {
			for _, item := range monkey.items {
				var worryLevel = monkey.operation(item)
				var throwTo = monkey.test(worryLevel)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
				monkey.items = monkey.items[1:]
				inspections[monkeyNum] += 1
			}
		}
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println(inspections)
	fmt.Println("Part 2: ", inspections[0]*inspections[1])
}

func parseItems(line string) []uint64 {
	var items = []uint64{}
	for _, item := range strings.Split(strings.TrimSpace(line), " ")[2:] {
		var itemValue, _ = strconv.Atoi(strings.ReplaceAll(item, ",", ""))
		items = append(items, uint64(itemValue))
	}
	return items
}

func parseOperation(line string) func(uint64) uint64 {
	var parts = strings.Split(strings.TrimSpace(line), " ")
	return func(old uint64) uint64 {
		var firstOperand, secondOperand = parseOperands(parts[3], parts[5], old)
		switch parts[4] {
		case "+":
			return firstOperand + secondOperand
		case "*":
			return firstOperand * secondOperand
		}
		panic("Unknown operator")
	}
}

func parseTest(lines []string) func(uint64) int {
	return func(worryLevel uint64) int {
		var operand, _ = strconv.Atoi(strings.Split(strings.TrimSpace(lines[0]), " ")[3])
		var trueCaseParts = strings.Split(strings.TrimSpace(lines[1]), " ")
		var falseCaseParts = strings.Split(strings.TrimSpace(lines[2]), " ")
		if worryLevel%uint64(operand) == 0 {
			var monkey, _ = strconv.Atoi(trueCaseParts[5])
			return monkey
		} else {
			var monkey, _ = strconv.Atoi(falseCaseParts[5])
			return monkey
		}
	}
}

func parseOperands(first, second string, old uint64) (uint64, uint64) {
	var firstIntValue, secondIntValue uint64
	if first == "old" {
		firstIntValue = old
	} else {
		var tmp, _ = strconv.Atoi(first)
		firstIntValue = uint64(tmp)
	}
	if second == "old" {
		secondIntValue = old
	} else {
		var tmp, _ = strconv.Atoi(second)
		secondIntValue = uint64(tmp)
	}
	return firstIntValue, secondIntValue
}
