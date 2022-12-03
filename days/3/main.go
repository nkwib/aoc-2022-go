package day_3

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func f(lines []string) int {
	prioritiesValue := 0
	for _, line := range lines {
		commonLetter := utils.CommonLetter(utils.SplitStringInHalf(line))
		val := utils.GetLetterValue(commonLetter)
		prioritiesValue += val
	}
	return prioritiesValue
}

func s(lines []string) int {
	prioritiesValue := 0
	// sliding window of 3
	for i := 0; i < len(lines)-2; i = i + 3 {
		line1, line2, line3 := lines[i], lines[i+1], lines[i+2]
		commonLetter := utils.CommonLetter(line1, line2, line3)
		val := utils.GetLetterValue(commonLetter)
		prioritiesValue += val
	}
	return prioritiesValue
}

func Solve(text string) {
	lines := strings.Split(text, "\n")
	total_1 := f(lines)
	total_2 := s(lines)

	fmt.Println("The sum of priorities of the items type is: ", total_1)
	fmt.Println("The sum of priorities for the badges is: ", total_2)
}
