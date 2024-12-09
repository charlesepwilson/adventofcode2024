package main

import (
	"advent_of_code_2024/day01"
	"advent_of_code_2024/day02"
	"advent_of_code_2024/day03"
	"advent_of_code_2024/day04"
	"advent_of_code_2024/day05"
	"advent_of_code_2024/day06"
	"advent_of_code_2024/day07"
	"advent_of_code_2024/day08"
	"advent_of_code_2024/utils"
	"slices"
)

func main() {
	solveDays := []int{
		8,
	}
	for _, solution := range []utils.DaySolution{
		day01.Solution{},
		day02.Solution{},
		day03.Solution{},
		day04.Solution{},
		day05.Solution{},
		day06.Solution{},
		day07.Solution{},
		day08.Solution{},
	} {
		if slices.Contains(solveDays, solution.Day()) {
			utils.Print(solution)
		}
	}
}
