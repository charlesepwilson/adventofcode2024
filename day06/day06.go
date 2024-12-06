package day06

import (
	"advent_of_code_2024/utils"
	"bytes"
	"slices"
)

const DAY = 6

type VectorI struct {
	down, right int
}

func (v VectorI) turnRight() VectorI {
	return VectorI{v.right, -v.down}
}

func (v VectorI) Add(other VectorI) VectorI {
	return VectorI{v.down + other.down, v.right + other.right}
}

func (v VectorI) Mul(val int) VectorI {
	return VectorI{v.down * val, v.right * val}
}

func parseInput(input []byte) (start VectorI, obstacles []VectorI, gridSize VectorI) {
	lines := bytes.Split(input, []byte("\n"))
	gridSize.down = len(lines)
	gridSize.right = len(lines[0])
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				obstacles = append(obstacles, VectorI{row, col})
			} else if char == '^' {
				start.down = row
				start.right = col
			}
		}
	}
	return start, obstacles, gridSize
}

func part1(input []byte) int {
	position, obstacles, gridSize := parseInput(input)
	facing := VectorI{down: -1, right: 0}
	pathPositions := utils.NewSet[VectorI]()
	for position.down >= 0 && position.right < gridSize.right && position.right >= 0 && position.down < gridSize.down {
		pathPositions.Add(position)
		nextPos := position.Add(facing)
		if slices.Contains(obstacles, nextPos) {
			facing = facing.turnRight()
		}
		nextPos = position.Add(facing)
		position = nextPos
	}

	return pathPositions.Len()
}

func part2(input []byte) int {
	//position, obstacles, gridSize := parseInput(input)
	//facing := VectorI{down: -1, right: 0}
	total := 0
	return total
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
