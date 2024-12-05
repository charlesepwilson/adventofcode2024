package day05

import (
	"advent_of_code_2024/utils"
	"bytes"
	"slices"
	"strconv"
)

const DAY = 5

func inOrder(pageNumbers []int, orderingFirst int, orderingLast int) bool {
	iFirst := slices.Index(pageNumbers, orderingFirst)
	if iFirst == -1 {
		return true
	}
	iLast := slices.Index(pageNumbers, orderingLast)
	if iLast == -1 {
		return true
	}
	return iFirst < iLast
}

type Ordering struct {
	first int
	last  int
}

func getFullOrder(orderings []Ordering) []int {
	var rCounts = make(map[int]int)
	for _, ordering := range orderings {
		rCounts[ordering.last] = rCounts[ordering.last] + 1
	}
	fullOrder := make([]int, len(rCounts))
	for i := 1; len(rCounts) > 0 && i < 10; i += 1 {
		for key, value := range rCounts {
			if value == i {
				fullOrder[i-1] = key
				delete(rCounts, key)
			}
		}
	}
	return fullOrder
}

func part1(input []byte) int {
	orderingPart, updatePart, _ := bytes.Cut(input, []byte("\n\n"))
	orderings := bytes.Split(orderingPart, []byte("\n"))
	total := 0
	for _, updateStr := range bytes.Split(updatePart, []byte("\n")) {
		pageNumberStrs := bytes.Split(updateStr, []byte(","))
		pageNumbers := make([]int, len(pageNumberStrs))
		for i, str := range pageNumberStrs {
			pageNumbers[i], _ = strconv.Atoi(string(str))
		}
		goodUpdate := true
		for _, ordering := range orderings {
			leftStr, rightStr, _ := bytes.Cut(ordering, []byte("|"))
			leftNum, _ := strconv.Atoi(string(leftStr))
			rightNum, _ := strconv.Atoi(string(rightStr))
			if !inOrder(pageNumbers, leftNum, rightNum) {
				goodUpdate = false
			}
		}
		if goodUpdate {
			total += pageNumbers[(len(pageNumbers)-1)/2]
		}

	}

	return total
}

func part2(input []byte) int {
	return 0
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
