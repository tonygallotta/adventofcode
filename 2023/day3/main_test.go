package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSpecialCharsRegex(t *testing.T) {
	pattern := regexp.MustCompile("[^\\d\\.]")
	fmt.Println(pattern.MatchString("467..114.."))
	fmt.Println("...*......"[3:4])
}

func TestGear(t *testing.T) {
	pattern := regexp.MustCompile("\\*")
	fmt.Println(pattern.FindAllStringIndex("467..*14.*", -1))
	fmt.Println("467..114.."[2:5])
}
