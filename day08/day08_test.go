package day08

import (
	"testing"
)

func getExample() []byte {
	return []byte("............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............")
}

func TestPart1(t *testing.T) {
	result := part1(getExample())
	if result != 14 {
		t.Errorf("Wrong answer for day %d p1: %d", DAY, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(getExample())
	if result != 11387 {
		t.Errorf("Wrong answer for day %d p2: %d", DAY, result)
	}
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
