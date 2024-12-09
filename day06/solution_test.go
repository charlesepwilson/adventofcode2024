package day06

import (
	"advent_of_code_2024/utils"
	"fmt"
	"sort"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPart1(Solution{}, t)
}

func TestPart2(t *testing.T) {
	utils.TestPart2(Solution{}, t)
}

func TestIsLoop(t *testing.T) {
	position, obstacles, gridSize := parseInput(Solution{}.GetExample())
	facing := utils.VectorI{Down: -1, Right: 0}
	rowObstacles := make([][]int, gridSize.Down)
	colObstacles := make([][]int, gridSize.Right)
	for _, obstacle := range obstacles {
		rowObstacles[obstacle.Down] = append(rowObstacles[obstacle.Down], obstacle.Right)
		colObstacles[obstacle.Right] = append(rowObstacles[obstacle.Right], obstacle.Down)
	}
	startState := State{position: position, facing: facing}
	if isLoop(
		startState,
		rowObstacles,
		colObstacles,
	) {
		t.Errorf("IsLoop returned true without extra obstacle")
	}

	loopMakers := []utils.VectorI{
		{Down: 6, Right: 3},
		{Down: 7, Right: 6},
		{Down: 7, Right: 7},
		{Down: 8, Right: 1},
		{Down: 8, Right: 3},
		{Down: 9, Right: 7},
	}
	for _, loopMaker := range loopMakers {
		fmt.Println("trialing ", loopMaker)
		trialRowObstacles := makeSliceCopy(rowObstacles)
		trialColObstacles := makeSliceCopy(colObstacles)
		trialRowObstacles[loopMaker.Down] = append(trialRowObstacles[loopMaker.Down], loopMaker.Right)
		sort.Ints(trialRowObstacles[loopMaker.Down])
		trialColObstacles[loopMaker.Right] = append(trialColObstacles[loopMaker.Right], loopMaker.Down)
		sort.Ints(trialColObstacles[loopMaker.Right])
		if !isLoop(
			startState,
			trialRowObstacles,
			trialColObstacles,
		) {
			t.Errorf("IsLoop returned false after adding the obstacle")
		}
	}

}
