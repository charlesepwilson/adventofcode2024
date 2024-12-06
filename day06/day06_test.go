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

func TestIsLoop(t *testing.T) {
	position, obstacles, gridSize := parseInput(getExample())
	facing := VectorI{down: -1, right: 0}
	rowObstacles := make([][]int, gridSize.down)
	colObstacles := make([][]int, gridSize.right)
	for _, obstacle := range obstacles {
		rowObstacles[obstacle.down] = append(rowObstacles[obstacle.down], obstacle.right)
		colObstacles[obstacle.right] = append(rowObstacles[obstacle.right], obstacle.down)
	}
	startState := State{position: position, facing: facing}
	if isLoop(
		startState,
		rowObstacles,
		colObstacles,
	) {
		t.Errorf("IsLoop returned true without extra obstacle")
	}
	newObstacle := VectorI{down: 6, right: 3}
	rowObstacles[newObstacle.down] = append(rowObstacles[newObstacle.down], newObstacle.right)
	colObstacles[newObstacle.right] = append(rowObstacles[newObstacle.right], newObstacle.down)

	if !isLoop(
		startState,
		rowObstacles,
		colObstacles,
	) {
		t.Errorf("IsLoop returned false after adding the obstacle")
	}

}
