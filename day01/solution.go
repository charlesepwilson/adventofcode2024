package day01

import (
	"advent_of_code_2024/utils"
	"bytes"
	"sort"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 1 }

func (Solution) Part1(input []byte) int {
	leftList, rightList := getLeftRightLists(input)
	sort.Ints(leftList)
	sort.Ints(rightList)

	result := 0
	for i := 0; i < len(leftList); i++ {
		result += utils.Abs(leftList[i] - rightList[i])
	}
	return result
}

func (Solution) Part2(input []byte) int {
	leftList, rightList := getLeftRightLists(input)
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

func (Solution) GetExample() []byte {
	return []byte("3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
}

func (Solution) ExampleAnswer1() int {
	return 11
}
func (Solution) ExampleAnswer2() int {
	return 31
}

func getLeftRightLists(input []byte) ([]int, []int) {
	lines := bytes.Split(input, []byte("\n"))
	var leftList []int
	var rightList []int

	for line := range lines {
		parts := bytes.Fields(lines[line])
		leftNum, _ := strconv.Atoi(string(parts[0]))
		rightNum, _ := strconv.Atoi(string(parts[1]))

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	return leftList, rightList
}
