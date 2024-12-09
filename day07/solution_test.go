package day07

import (
	"advent_of_code_2024/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPart1(Solution{}, t)
}

func TestPart2(t *testing.T) {
	utils.TestPart2(Solution{}, t)
}

func TestCountDigits(t *testing.T) {
	cases := [][]int{
		{1, 1},
		{2, 1},
		{10, 2},
		{11, 2},
		{99, 2},
		{100, 3},
		{999, 3},
	}
	for _, c := range cases {
		result := countDigits(c[0])
		if result != c[1] {
			t.Errorf("Count of digits for %d recorded as %d", c[0], result)
		}
	}

}
