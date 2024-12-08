package day07

import (
	"testing"
)

func getExample() []byte {
	return []byte("190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20")
}

func TestPart1(t *testing.T) {
	result := part1(getExample())
	if result != 3749 {
		t.Errorf("Wrong answer for day %d p1: %d", DAY, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(getExample())
	if result != 6 {
		t.Errorf("Wrong answer for day %d p2: %d", DAY, result)
	}
}
