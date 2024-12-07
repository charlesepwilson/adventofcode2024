package day06

import (
	"advent_of_code_2024/utils"
	"bytes"
	"fmt"
	"slices"
	"sort"
)

const DAY = 6

type VectorI struct {
	down, right int
}

// if we make a turn against an obstacle h0 at (r0, c0) while moving up
// then the state s0 at the turn point is (r0+1, c0, U)
// and there must be an obstacle h1 at (r0+1, c1), where c1 > c0
// then the next turn state s1 is (r0+1, c1-1, R)
// then next obstacle h2 is at (r2, c1-1), where r2 > r0+1
// and next state s2 is (r2 + 1, c1-1, D)

// we want to find a sequence of states such that for some value n, sn == s0
// since every turn goes 90 degrees we know that the size of the loop must be a multiple of 4
// we know that every obstacle in the loop (r, c) must have a previous obstacle that's offset by one row/column and a next obstcle also offset by one row or column

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

type State struct {
	position, facing VectorI
}

func (s State) Valid() bool {
	return s.position.down >= 0 && s.position.right >= 0
}

func findNextPos(state State, rowObstacles [][]int, colObstacles [][]int) (nextState State) {
	var relevantObstacles [][]int
	var alignment int
	var displacement int
	var fVal int
	sideways := state.facing.down == 0
	if sideways {
		relevantObstacles = rowObstacles
		alignment = state.position.down
		displacement = state.position.right
		fVal = state.facing.right
	} else {
		relevantObstacles = colObstacles
		alignment = state.position.right
		displacement = state.position.down
		fVal = state.facing.down
	}
	alignedObstacles := relevantObstacles[alignment]
	backwards := fVal < 0
	if backwards {
		slices.Reverse(alignedObstacles)
	}
	for _, obstaclePos := range alignedObstacles {
		diff := obstaclePos < displacement
		if diff == backwards {
			next := obstaclePos - fVal
			if sideways {
				nextState.position.right = next
				nextState.position.down = alignment
			} else {
				nextState.position.right = alignment
				nextState.position.down = next
			}
			nextState.facing = state.facing.turnRight()
			//fmt.Println("state", state, "goes to state", nextState)
			return nextState

		}
	}
	return State{VectorI{-1, -1}, state.facing}
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

func part2(input []byte) int {
	startPos, obstacles, gridSize := parseInput(input)
	facing := VectorI{down: -1, right: 0}
	rowObstacles := make([][]int, gridSize.down)
	colObstacles := make([][]int, gridSize.right)
	for _, obstacle := range obstacles {
		rowObstacles[obstacle.down] = append(rowObstacles[obstacle.down], obstacle.right)
		colObstacles[obstacle.right] = append(colObstacles[obstacle.right], obstacle.down)
	}
	startState := State{startPos, facing}
	total := 0
	fmt.Println(rowObstacles)
	fmt.Println(colObstacles)
	loopMakers := make([]VectorI, 0)
	for i := range gridSize.right {
		for j := range gridSize.down {
			trialObstacle := VectorI{down: j, right: i}
			if trialObstacle == startPos {
				continue
			}
			trialRowObstacles := makeSliceCopy(rowObstacles)
			trialColObstacles := makeSliceCopy(colObstacles)
			trialRowObstacles[j] = append(trialRowObstacles[j], i)
			sort.Ints(trialRowObstacles[j])
			trialColObstacles[i] = append(trialColObstacles[i], j)
			sort.Ints(trialColObstacles[i])
			if isLoop(startState, trialRowObstacles, trialColObstacles) {
				total += 1
				loopMakers = append(loopMakers, trialObstacle)
				//fmt.Println(i, j)
			}
		}
	}
	//fmt.Println(loopMakers)
	return total
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
