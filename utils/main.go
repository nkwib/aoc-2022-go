package utils

import (
	"fmt"
	"os"
	"strings"
)

// export sumInts
func SumInts(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text := string(data)
	return text
}

func SplitStringInHalf(text string) (string, string) {
	middle := len(text) / 2
	return text[:middle], text[middle:]
}

func CommonLetter(lists ...string) rune {
	var commonBetweenLists rune
	firstList := lists[0]
	for _, letter := range firstList {
		common := true
		for _, list := range lists[1:] {
			if !strings.ContainsRune(list, letter) {
				common = false
				break
			}
		}
		if common {
			commonBetweenLists = letter
			break
		}
	}
	return commonBetweenLists
}

func GetLetterValue(letter rune) int {
	val := int(letter)
	if val > 96 {
		return val - 96
	}
	return val - 38
}
