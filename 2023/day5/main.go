package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day5/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

type AlmanacMapFunction struct {
	dest   int
	source int
	length int
}

type AlmanacMap struct {
	label string
	maps  []AlmanacMapFunction
}

func (a AlmanacMapFunction) HasMapping(num int) bool {
	var diff = num - a.source
	return diff >= 0 && diff < a.length
}

func (a AlmanacMapFunction) MapValue(num int) int {
	var diff = num - a.source
	if diff >= 0 && diff < a.length {
		return a.dest + diff
	}
	return num
}

func (a AlmanacMap) MapValue(num int) int {
	for _, mapFn := range a.maps {
		if mapFn.HasMapping(num) {
			//fmt.Printf("%v has a mapping for %v\n", a.label, num)
			return mapFn.MapValue(num)
		}
	}
	//fmt.Printf("%v does NOT have a mapping for %v\n", a.label, num)
	return num
}
func part1(lines []string) {
	var answer = 0
	var seeds, maps = ParseInputs(lines)
	for i, seed := range seeds {
		for _, m := range maps {
			seed = m.MapValue(seed)
		}
		if seed < answer || i == 0 {
			answer = seed
		}
	}
	fmt.Println("Part 1: ", answer)
}

func part2(lines []string) {
	var answer = -1
	var seeds, maps = ParseInputs(lines)
	for i := 0; i < len(seeds)-1; i += 2 {
		// Well, there are way too many seeds to solve it this way. Maybe we can start with locations and work backward?
		fmt.Printf("Testings seeds %v-%v\n", seeds[i], seeds[i]+seeds[i+1]-1)
		var locationNumber = 0
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			locationNumber = seed
			for _, m := range maps {
				locationNumber = m.MapValue(locationNumber)
			}
			if locationNumber < answer || answer == -1 {
				answer = locationNumber
			}
		}
	}
	fmt.Println("Part 2: ", answer)
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

func ParseInputs(lines []string) ([]int, []AlmanacMap) {
	numberPattern := regexp.MustCompile(`\d+`)
	var seeds []int
	var mapFunctions = []AlmanacMapFunction{}
	var maps = []AlmanacMap{}
	var label = ""
	for i, line := range lines {
		matchedNumbers := MapToInts(numberPattern.FindAllString(line, -1))
		if i == 0 {
			seeds = matchedNumbers
		} else if len(matchedNumbers) > 0 {
			mapFunctions = append(mapFunctions, AlmanacMapFunction{matchedNumbers[0], matchedNumbers[1], matchedNumbers[2]})
		} else if len(line) > 0 {
			label, _ = SplitToPair(line, " ")
		} else if i > 3 {
			maps = append(maps, AlmanacMap{label, mapFunctions})
			mapFunctions = []AlmanacMapFunction{}
		}
	}
	maps = append(maps, AlmanacMap{label, mapFunctions})
	return seeds, maps
}
