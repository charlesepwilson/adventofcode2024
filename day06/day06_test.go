package day06

import (
	"testing"
)

func getExample() []byte {
	return []byte("....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...")
}

func TestPart1(t *testing.T) {
	result := part1(getExample())
	if result != 41 {
		t.Errorf("Wrong answer for day %d p1: %d", DAY, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(getExample())
	if result != 6 {
		t.Errorf("Wrong answer for day %d p2: %d", DAY, result)
	}
}
