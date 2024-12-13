package day10

import (
	"advent_of_code_2024/utils"
)

type Solution struct{}

func (Solution) Day() int { return 10 }

func (Solution) Part1(input []byte) int {
	grid := utils.GetIntegerGrid(input)
	trails := findAllTrailEnds(grid)
	result := 0
	for _, ends := range trails {
		result += ends.Len()
	}
	return result
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample() []byte {
	return []byte("89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732")
}

func (Solution) ExampleAnswer1() int {
	return 36
}
func (Solution) ExampleAnswer2() int {
	return 0
}

const startHeight = 0
const endHeight = 9

func buildNetwork(grid [][]int) (map[utils.VectorI]utils.Set[utils.VectorI], []utils.VectorI, []utils.VectorI) {
	network := make(map[utils.VectorI]utils.Set[utils.VectorI])
	gridSize := utils.VectorI{Down: len(grid), Right: len(grid[0])}
	trailHeads := make([]utils.VectorI, 0)
	trailEnds := make([]utils.VectorI, 0)
	for i, line := range grid {
		for j, v := range line {
			position := utils.VectorI{Down: i, Right: j}
			if v == startHeight {
				trailHeads = append(trailHeads, position)
			} else if v == endHeight {
				trailEnds = append(trailEnds, position)
			}
			for _, adjacent := range position.GetCardinalAdjacents() {
				if utils.WithinGrid(adjacent, gridSize) && (grid[adjacent.Down][adjacent.Right]-v) == 1 {
					_, ok := network[position]
					if !ok {
						network[position] = utils.NewSet[utils.VectorI]()
					}
					network[position].Add(adjacent)
				}
			}
		}
	}
	return network, trailHeads, trailEnds
}

func findTrailEnds(trailhead utils.VectorI, network map[utils.VectorI]utils.Set[utils.VectorI]) utils.Set[utils.VectorI] {
	positions := utils.NewSet[utils.VectorI]()
	positions.Add(trailhead)
	for i := 0; i < endHeight; i++ {
		newPositions := utils.NewSet[utils.VectorI]()
		for position := range positions.Iterate() {
			nextPositions := network[position]
			for nextPosition := range nextPositions.Iterate() {
				newPositions.Add(nextPosition)
			}
		}
		positions = newPositions
	}
	return positions
}

func findAllTrailEnds(grid [][]int) map[utils.VectorI]utils.Set[utils.VectorI] {
	network, trailheads, _ := buildNetwork(grid)
	trails := make(map[utils.VectorI]utils.Set[utils.VectorI])
	for _, head := range trailheads {
		trails[head] = findTrailEnds(head, network)
	}
	return trails
}
