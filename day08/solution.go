package day08

import (
	"advent_of_code_2024/utils"
	"bytes"
)

type Solution struct{}

func (Solution) Day() int { return 8 }

func (Solution) Part1(input []byte) int {
	return countAntiNodes(input, part1NodeFinder)
}

func (Solution) Part2(input []byte) int {
	return countAntiNodes(input, part2NodeFinder)
}

func (Solution) GetExample(part int) []byte {
	return []byte("............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............")
}

func (Solution) ExampleAnswer1() int {
	return 14
}
func (Solution) ExampleAnswer2() int {
	return 34
}

func parseInput(input []byte) (antennae map[byte][]utils.VectorI, gridSize utils.VectorI) {
	antennae = make(map[byte][]utils.VectorI)
	lines := bytes.Split(input, []byte("\n"))
	gridSize.Down = len(lines)
	gridSize.Right = len(lines[0])
	background := byte('.')
	for row, line := range lines {
		for col, v := range line {
			if v != background {
				antennae[v] = append(antennae[v], utils.VectorI{Down: row, Right: col})
			}
		}
	}
	return antennae, gridSize
}

func getPairs(locations []utils.VectorI) [][]utils.VectorI {
	n := len(locations)
	numPairs := n * (n - 1) / 2
	pairs := make([][]utils.VectorI, numPairs)
	pairIndex := 0
	for i := 0; i < n; i++ {
		//fmt.Println("i", i)
		for j := i + 1; j < n; j++ {
			//fmt.Println("j", j)

			p := []utils.VectorI{locations[i], locations[j]}
			pairs[pairIndex] = p
			pairIndex++
		}

	}
	return pairs
}

type NodeFinder func(pair []utils.VectorI, gridSize utils.VectorI) utils.Set[utils.VectorI]

func part1NodeFinder(pair []utils.VectorI, gridSize utils.VectorI) utils.Set[utils.VectorI] {
	antiNodes := utils.NewSet[utils.VectorI]()
	diff := pair[0].Sub(pair[1])
	up := pair[0].Add(diff)
	down := pair[1].Sub(diff)
	for _, aNode := range []utils.VectorI{up, down} {
		if utils.WithinGrid(aNode, gridSize) {
			antiNodes.Add(aNode)
		}
	}
	return antiNodes
}

func part2NodeFinder(pair []utils.VectorI, gridSize utils.VectorI) utils.Set[utils.VectorI] {
	antiNodes := utils.NewSet[utils.VectorI]()
	diff := pair[0].Sub(pair[1])
	diff = diff.Simplify()
	aNode := pair[0]
	for utils.WithinGrid(aNode, gridSize) {
		antiNodes.Add(aNode)
		aNode = aNode.Add(diff)
	}
	aNode = pair[0].Sub(diff)
	for utils.WithinGrid(aNode, gridSize) {
		antiNodes.Add(aNode)
		aNode = aNode.Sub(diff)
	}
	return antiNodes
}

func countAntiNodes(input []byte, nodeFinder NodeFinder) int {
	antennae, gridSize := parseInput(input)
	antiNodes := utils.NewSet[utils.VectorI]()

	for _, locations := range antennae {
		//fmt.Println(locations)
		for _, pair := range getPairs(locations) {
			//fmt.Println("pair", pair)
			for antiNode := range nodeFinder(pair, gridSize).Iterate() {
				antiNodes.Add(antiNode)
			}
		}
	}

	return antiNodes.Len()
}
