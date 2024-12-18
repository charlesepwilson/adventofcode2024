package day17

import (
	"advent_of_code_2024/utils"
	"fmt"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPart1(Solution{}, t)
}

func TestPart2(t *testing.T) {
	utils.TestPart2(Solution{}, t)
}

func TestCandidates(t *testing.T) {
	program := []int{2, 4, 1, 2, 7, 5, 1, 3, 4, 3, 5, 5, 0, 3, 3, 0}
	a := cheatPart2(program)
	fmt.Println(a)
	abin := strconv.FormatInt(int64(a), 2)
	fmt.Println(abin, len(abin))
	if a>>48 != 0 {
		t.Error("too big", a, abin, a>>48)
	}
	if a>>45 == 0 {
		t.Error("too small", a, abin, a>>45)
	}
	for a > 0 {
		fmt.Println(computeOutput(a))
		a = a >> 3
	}
}
