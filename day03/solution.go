package day03

import (
	"regexp"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 3 }

func (Solution) Part1(input []byte) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	result := 0
	for _, match := range re.FindAllSubmatch(input, -1) {
		left, _ := strconv.Atoi(string(match[1]))
		right, _ := strconv.Atoi(string(match[2]))
		result += left * right
	}
	return result
}

func (Solution) Part2(input []byte) int {
	return totalWithDoDont(input)
}

func (Solution) GetExample(part int) []byte {
	return []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
}

func (Solution) ExampleAnswer1() int {
	return 161
}
func (Solution) ExampleAnswer2() int {
	return 48
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
