package day_2

import (
	"fmt"
	"strings"
)

func Solve(text string) {
	mappedVals := map[string][]int{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	games := strings.Split(text, "\n")
	values := make([][]int, len(games))
	for index, game := range games {
		values[index] = mappedVals[game]
	}
	total_1, total_2 := 0, 0
	for _, val := range values {
		total_1 += val[0]
		total_2 += val[1]
	}

	fmt.Println("Playing your interpretation of Rock, Paper, Scissors will give you: ", total_1)
	fmt.Println("Playing as intended by the Elf at Rock, Paper, Scissors will give you: ", total_2)
}
