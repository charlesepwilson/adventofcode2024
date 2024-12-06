package day05

import (
	"advent_of_code_2024/utils"
	"bytes"
	"slices"
	"strconv"
)

const DAY = 5

func inOrder(pageNumbers []int, beforeRules map[int][]int) bool {
	seen := utils.NewSet[int]()
	for _, n := range pageNumbers {
		for _, mustBeBefore := range beforeRules[n] {
			if seen.Contains(mustBeBefore) {
				return false
			}
		}
		seen.Add(n)
	}
	return true
}

type Ordering struct {
	first int
	last  int
}

func orderingFromBytes(text []byte) Ordering {
	leftStr, rightStr, _ := bytes.Cut(text, []byte("|"))
	leftNum, _ := strconv.Atoi(string(leftStr))
	rightNum, _ := strconv.Atoi(string(rightStr))
	return Ordering{leftNum, rightNum}
}

func getMiddleValue(nums []int) int {
	return nums[(len(nums)-1)/2]
}

func parseInput(input []byte) (rules map[int][]int, updates [][]int) {
	orderingPart, updatePart, _ := bytes.Cut(input, []byte("\n\n"))
	orderingStrs := bytes.Split(orderingPart, []byte("\n"))
	beforeRules := make(map[int][]int)
	for _, ordering := range orderingStrs {
		o := orderingFromBytes(ordering)
		beforeRules[o.first] = append(beforeRules[o.first], o.last)
	}
	updateLines := bytes.Split(updatePart, []byte("\n"))
	updateList := make([][]int, len(updateLines))
	for j, updateStr := range updateLines {
		pageNumberStrs := bytes.Split(updateStr, []byte(","))
		pageNumbers := make([]int, len(pageNumberStrs))
		for i, str := range pageNumberStrs {
			pageNumbers[i], _ = strconv.Atoi(string(str))
		}
		updateList[j] = pageNumbers
	}
	return beforeRules, updateList
}

func part1(input []byte) int {
	rules, updates := parseInput(input)
	total := 0
	for _, pageNumbers := range updates {
		if inOrder(pageNumbers, rules) {
			total += getMiddleValue(pageNumbers)
		}
	}
	return total
}

func part2(input []byte) int {
	beforeRules, updates := parseInput(input)
	total := 0
	for _, pageNumbers := range updates {
		if !inOrder(pageNumbers, beforeRules) {
			slices.SortFunc(
				pageNumbers,
				func(i int, j int) int {
					if slices.Contains(beforeRules[i], j) {
						return -1
					} else if slices.Contains(beforeRules[j], i) {
						return 1
					} else {
						return 0
					}
				},
			)
			total += getMiddleValue(pageNumbers)
		}
	}
	return total
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
