// print hello world
package main

import (
	"aoc/days/1"
	"aoc/days/2"
	"aoc/utils"
)

func readFile(filename string) string {
	text := utils.ReadFile(filename)
	return text
}

func main() {
	day_1.Solve(readFile("input_one.txt"))
	day_2.Solve(readFile("input_two.txt"))
}
