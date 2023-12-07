package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	data, _ := os.ReadFile("day5/input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines) // Sample answer = 46
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

type Range struct {
	min int
	max int
}

func (r1 Range) Overlaps(r2 Range) bool {
	return (r1.min >= r2.min && r1.min < r2.max) || (r1.max >= r2.min && r1.min < r2.max)
}

func (a AlmanacMapFunction) HasSourceMapping(num int) bool {
	var diff = num - a.source
	return diff >= 0 && diff < a.length
}

func (a AlmanacMapFunction) HasDestOverlap(r Range) bool {
	return a.DestRange().Overlaps(r)
}

func (a AlmanacMapFunction) DestRange() Range {
	return Range{a.dest, a.dest + a.length - 1}
}

func (a AlmanacMapFunction) SourceRange() Range {
	return Range{a.dest, a.dest + a.length - 1}
}

func (a AlmanacMapFunction) MapSourceValue(num int) int {
	var diff = num - a.source
	if diff >= 0 && diff < a.length {
		return a.dest + diff
	}
	return num
}

func (a AlmanacMap) MapValue(num int) int {
	for _, mapFn := range a.maps {
		if mapFn.HasSourceMapping(num) {
			//fmt.Printf("%v has a mapping for %v\n", a.label, num)
			return mapFn.MapSourceValue(num)
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
	var seedRanges = []Range{}
	for i := 0; i < len(seeds)-1; i += 2 {
		seedRanges = append(seedRanges, Range{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	var mins = []int{}
	var wg sync.WaitGroup
	for i, s := range seedRanges {
		wg.Add(1)
		go func(r Range, rangeNum int) {
			defer wg.Done()
			fmt.Println("Checking seed range ", rangeNum)
			var minForRange = MinInRange(r, maps)
			mins = append(mins, minForRange)
			fmt.Println("Min for seed range", r, minForRange)
		}(s, i)
	}
	wg.Wait()
	answer = mins[0]
	for _, m := range mins {
		if m < answer {
			answer = m
		}
	}
	// 368703957 too high
	// 2494335141
	// 207836694
	// 95461669
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

func MinInRange(s Range, maps []AlmanacMap) int {
	var answer = math.MaxInt32
	for seed := s.min; seed <= s.max; seed++ {
		v := seed
		for _, m := range maps {
			v = m.MapValue(v)
		}
		if v < answer {
			answer = v
		}
	}
	return answer
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
