package day16

import (
	"advent_of_code_2024/utils"
	"testing"
)

func GetExample2() []byte {
	return []byte("#################\n#...#...#...#..E#\n#.#.#.#.#.#.#.#.#\n#.#.#.#...#...#.#\n#.#.#.#.###.#.#.#\n#...#.#.#.....#.#\n#.#.#.#.#.#####.#\n#.#...#.#.#.....#\n#.#.#####.#.###.#\n#.#.#.......#...#\n#.#.###.#####.###\n#.#.#...#.....#.#\n#.#.#.#####.###.#\n#.#.#.........#.#\n#.#.#.#########.#\n#S#.............#\n#################")
}

func TestPart1(t *testing.T) {
	utils.TestPart1(Solution{}, t)
}

func TestPart2(t *testing.T) {
	utils.TestPart2(Solution{}, t)
}

func TestPart1Ex2(t *testing.T) {
	day := Solution{}
	exampleInput := GetExample2()
	result := day.Part1(exampleInput)
	if result != 11048 {
		t.Errorf("Wrong answer for day %d p1: %d", day.Day(), result)
	}
}

func TestPart2Ex2(t *testing.T) {
	day := Solution{}
	exampleInput := GetExample2()
	result := day.Part2(exampleInput)
	if result != 45 {
		t.Errorf("Wrong answer for day %d p2: %d", day.Day(), result)
	}
}
