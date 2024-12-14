package day05

import (
	"advent_of_code_2024/utils"
	"bytes"
	"slices"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 5 }

func (Solution) Part1(input []byte) int {
	rules, updates := parseInput(input)
	total := 0
	for _, pageNumbers := range updates {
		if inOrder(pageNumbers, rules) {
			total += getMiddleValue(pageNumbers)
		}
	}
	return total
}

func (Solution) Part2(input []byte) int {
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

func (Solution) GetExample(part int) []byte {
	return []byte("47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
}

func (Solution) ExampleAnswer1() int {
	return 143
}
func (Solution) ExampleAnswer2() int {
	return 123
}

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
