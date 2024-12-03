package day03

import (
	"advent_of_code_2024/utils"
	"regexp"
	"strconv"
)

const DAY = 3

func Part1() int {
	input := utils.ReadInput(DAY, 1)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	result := 0
	for _, match := range re.FindAllSubmatch(input, -1) {
		left, _ := strconv.Atoi(string(match[1]))
		right, _ := strconv.Atoi(string(match[2]))
		result += left * right
	}
	return result
}

func totalWithDoDont(text []byte) int {
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	active := true
	result := 0
	for _, match := range re.FindAllSubmatch(text, -1) {
		if len(match[1]) > 0 {
			if active {
				left, _ := strconv.Atoi(string(match[2]))
				right, _ := strconv.Atoi(string(match[3]))
				result += left * right
			}
		} else if len(match[4]) > 0 {
			active = true
		} else if len(match[5]) > 0 {
			active = false
		}

	}
	return result
}

func Part2() int {
	input := utils.ReadInput(DAY, 1)
	return totalWithDoDont(input)
}
