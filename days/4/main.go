package day_4

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func f(lines [][][]int) int {
	overlaps := 0
	for _, line := range lines {
		if utils.OneContainsTheOther(line[0], line[1]) {
			overlaps++
		}
	}
	return overlaps
}

func s(lines [][][]int) int {
	inclusions := 0
	for _, line := range lines {
		if utils.OneOverlapsTheOther(line[0], line[1]) {
			inclusions++
		}
	}
	return inclusions
}

func Solve(text string) {
	lines := strings.Split(text, "\n")
	pairs := make([][][]int, len(lines))
	for index, line := range lines {
		var pair [][]int = make([][]int, 2)
		for i, pairText := range strings.Split(line, ",") {
			pair[i] = StringToIntSlice(pairText, "-")
		}
		pairs[index] = pair
	}
	fmt.Println("The sum of priorities of the items type is: ", f(pairs))
	fmt.Println("The sum of priorities for the badges is: ", s(pairs))
}

func StringToIntSlice(text, sep string) []int {
	slice := make([]int, 2)
	for i, val := range strings.Split(text, sep) {
		slice[i], _ = strconv.Atoi(val)
	}
	return slice
}
