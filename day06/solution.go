package day06

import (
	"advent_of_code_2024/utils"
	"bytes"
	"fmt"
	"slices"
	"sort"
)

type Solution struct{}

func (Solution) Day() int { return 6 }

func (Solution) Part1(input []byte) int {
	position, obstacles, gridSize := parseInput(input)
	pathPositions := getPathPositions(position, obstacles, gridSize)
	return pathPositions.Len()
}

func (Solution) Part2(input []byte) int {
	startPos, obstacles, gridSize := parseInput(input)
	facing := utils.VectorI{Down: -1, Right: 0}
	rowObstacles := make([][]int, gridSize.Down)
	colObstacles := make([][]int, gridSize.Right)
	for _, obstacle := range obstacles {
		rowObstacles[obstacle.Down] = append(rowObstacles[obstacle.Down], obstacle.Right)
		colObstacles[obstacle.Right] = append(colObstacles[obstacle.Right], obstacle.Down)
	}
	for _, oblist := range rowObstacles {
		sort.Ints(oblist)
	}
	for _, oblist := range colObstacles {
		sort.Ints(oblist)
	}
	startState := State{startPos, facing}
	total := 0
	fmt.Println(obstacles)
	fmt.Println(rowObstacles)
	fmt.Println(colObstacles)
	loopMakers := make([]utils.VectorI, 0)
	pathPositions := getPathPositions(startPos, obstacles, gridSize)
	for trialObstacle := range pathPositions.Iterate() {
		if trialObstacle == startPos {
			continue
		}
		trialRowObstacles := makeSliceCopy(rowObstacles)
		trialColObstacles := makeSliceCopy(colObstacles)
		trialRowObstacles[trialObstacle.Down] = append(trialRowObstacles[trialObstacle.Down], trialObstacle.Right)
		sort.Ints(trialRowObstacles[trialObstacle.Down])
		trialColObstacles[trialObstacle.Right] = append(trialColObstacles[trialObstacle.Right], trialObstacle.Down)
		sort.Ints(trialColObstacles[trialObstacle.Right])
		if isLoop(startState, trialRowObstacles, trialColObstacles) {
			total += 1
			loopMakers = append(loopMakers, trialObstacle)
		}
	}
	//fmt.Println(loopMakers)
	return total
}

func (Solution) GetExample(part int) []byte {
	return []byte("....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...")
}

func (Solution) ExampleAnswer1() int {
	return 41
}
func (Solution) ExampleAnswer2() int {
	return 6
}

func parseInput(input []byte) (start utils.VectorI, obstacles []utils.VectorI, gridSize utils.VectorI) {
	lines := bytes.Split(input, []byte("\n"))
	gridSize.Down = len(lines)
	gridSize.Right = len(lines[0])
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				obstacles = append(obstacles, utils.VectorI{Down: row, Right: col})
			} else if char == '^' {
				start.Down = row
				start.Right = col
			}
		}
	}
	return start, obstacles, gridSize
}

func getPathPositions(position utils.VectorI, obstacles []utils.VectorI, gridSize utils.VectorI) utils.Set[utils.VectorI] {
	facing := utils.VectorI{Down: -1, Right: 0}
	pathPositions := utils.NewSet[utils.VectorI]()
	for position.Down >= 0 && position.Right < gridSize.Right && position.Right >= 0 && position.Down < gridSize.Down {
		pathPositions.Add(position)
		nextPos := position.Add(facing)
		if slices.Contains(obstacles, nextPos) {
			facing = facing.TurnRight()
		}
		nextPos = position.Add(facing)
		position = nextPos
	}
	return pathPositions
}

type State struct {
	position, facing utils.VectorI
}

func (s State) Valid() bool {
	return s.position.Down >= 0 && s.position.Right >= 0
}

func findNextPos(state State, rowObstacles [][]int, colObstacles [][]int) (nextState State) {
	var relevantObstacles [][]int
	var alignment int
	var displacement int
	var fVal int
	sideways := state.facing.Down == 0
	if sideways {
		relevantObstacles = rowObstacles
		alignment = state.position.Down
		displacement = state.position.Right
		fVal = state.facing.Right
	} else {
		relevantObstacles = colObstacles
		alignment = state.position.Right
		displacement = state.position.Down
		fVal = state.facing.Down
	}
	alignedObstacles := relevantObstacles[alignment]
	backwards := fVal < 0
	if backwards {
		copyObstacles := make([]int, len(alignedObstacles))
		copy(copyObstacles, alignedObstacles)
		slices.Reverse(copyObstacles)
		alignedObstacles = copyObstacles
	}
	for _, obstaclePos := range alignedObstacles {
		diff := obstaclePos < displacement
		if diff == backwards {
			next := obstaclePos - fVal
			if sideways {
				nextState.position.Right = next
				nextState.position.Down = alignment
			} else {
				nextState.position.Right = alignment
				nextState.position.Down = next
			}
			nextState.facing = state.facing.TurnRight()
			//fmt.Println("state", state, "goes to state", nextState)
			return nextState

		}
	}
	return State{utils.VectorI{Down: -1, Right: -1}, state.facing}
}

func isLoop(state State, rowObstacles [][]int, colObstacles [][]int) bool {
	seen := utils.NewSet[State]()
	path := make([]State, 0)
	for {
		if !state.Valid() {
			//for _, st := range path {
			//	fmt.Println(st)
			//}
			return false
		} else if seen.Contains(state) {
			//fmt.Println("Found loop")
			//for _, st := range path {
			//	fmt.Println(st)
			//}
			return true
		}
		seen.Add(state)
		path = append(path, state)
		state = findNextPos(state, rowObstacles, colObstacles)
	}
}

func makeSliceCopy(slice [][]int) [][]int {
	newSlice := make([][]int, len(slice))
	for i := range newSlice {
		newSlice[i] = make([]int, len(slice[i]))
		copy(newSlice[i], slice[i])
	}
	return newSlice
}
