package solutions

import (
	"advent_of_code_2024/utils"
	"bytes"
	"sort"
	"strconv"
)

const DAY = 1

func getLeftRightLists() ([]int, []int) {
	input := utils.ReadInput(DAY, 1)
	lines := bytes.Split(input, []byte("\n"))
	var leftList []int
	var rightList []int

	for line := range lines {
		parts := bytes.Fields(lines[line])
		leftNum, _ := strconv.ParseInt(string(parts[0]), 10, 0)
		rightNum, _ := strconv.ParseInt(string(parts[1]), 10, 0)

		leftList = append(leftList, int(leftNum))
		rightList = append(rightList, int(rightNum))
	}
	return leftList, rightList
}

func Part1() int {
	leftList, rightList := getLeftRightLists()
	sort.Ints(leftList)
	sort.Ints(rightList)

	result := 0
	for i := 0; i < len(leftList); i++ {
		result += utils.Abs(leftList[i] - rightList[i])
	}
	return result
}

func Part2() int {
	leftList, rightList := getLeftRightLists()
	counts := make(map[int]int)
	result := 0
	for _, item := range leftList {
		count, ok := counts[item]
		if !ok {
			counts[item] = utils.Count(rightList, item)
			count = counts[item]
		}
		result += item * count
	}
	return result
}
