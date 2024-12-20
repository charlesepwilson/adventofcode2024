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
	"advent_of_code_2024/day09"
	"advent_of_code_2024/day10"
	"advent_of_code_2024/day11"
	"advent_of_code_2024/day12"
	"advent_of_code_2024/day13"
	"advent_of_code_2024/day14"
	"advent_of_code_2024/day15"
	"advent_of_code_2024/day16"
	"advent_of_code_2024/day17"
	"advent_of_code_2024/day18"
	"advent_of_code_2024/day19"
	"advent_of_code_2024/utils"
	"slices"
)

func main() {
	solveDays := []int{
		19,
	}
	for _, solution := range []utils.DaySolution[int]{
		day01.Solution{},
		day02.Solution{},
		day03.Solution{},
		day04.Solution{},
		day05.Solution{},
		day06.Solution{},
		day07.Solution{},
		day08.Solution{},
		day09.Solution{},
		day10.Solution{},
		day11.Solution{},
		day12.Solution{},
		day13.Solution{},
		day14.Solution{},
		day15.Solution{},
		day16.Solution{},
		day19.Solution{},
	} {
		if slices.Contains(solveDays, solution.Day()) {
			utils.Print(solution)
		}
	}
	for _, solution := range []utils.DaySolution[string]{
		day17.Solution{},
		day18.Solution{},
	} {
		if slices.Contains(solveDays, solution.Day()) {
			utils.Print(solution)
		}
	}
}
