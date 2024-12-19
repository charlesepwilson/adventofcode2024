package day18

import (
	"advent_of_code_2024/utils"
	"bytes"
	"fmt"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 18 }

func (Solution) Part1(input []byte) string {
	vectors := parseInput(input)[:numBlocks]
	start := utils.VectorI{}
	end := gridSize.Sub(diagVector)
	vSet := utils.SetFromSlice(vectors)
	distance, _ := dijkstra(start, end, vSet)
	return strconv.Itoa(distance)
}

func (Solution) Part2(input []byte) string {
	vectors := parseInput(input)
	start := utils.VectorI{}
	end := gridSize.Sub(diagVector)
	vSet := utils.NewSet[utils.VectorI]()
	var answer utils.VectorI
	for _, v := range vectors {
		vSet.Add(v)
		_, success := dijkstra(start, end, vSet)
		if !success {
			answer = v
			break
		}
	}
	return fmt.Sprintf("%d,%d", answer.Right, answer.Down)
}

func (Solution) GetExample(part int) []byte {
	gridSize = diagVector.Mul(7)
	numBlocks = 12
	return []byte("5,4\n4,2\n4,5\n3,0\n2,1\n6,3\n2,4\n1,5\n0,6\n3,3\n2,6\n5,1\n1,2\n5,5\n2,5\n6,5\n1,4\n0,4\n6,4\n1,1\n6,1\n1,0\n0,5\n1,6\n2,0")
}

func (Solution) ExampleAnswer1() string {
	return "22"
}
func (Solution) ExampleAnswer2() string {
	return "6,1"
}

var diagVector = utils.VectorI{Right: 1, Down: 1}
var gridSize = diagVector.Mul(71)
var numBlocks = 1024

func parseInput(input []byte) []utils.VectorI {
	lines := bytes.Split(input, []byte("\n"))
	vectors := make([]utils.VectorI, len(lines))
	for i, line := range lines {
		rightB, downB, _ := bytes.Cut(line, []byte(","))
		right, _ := strconv.Atoi(string(rightB))
		down, _ := strconv.Atoi(string(downB))
		vectors[i] = utils.VectorI{Right: right, Down: down}
	}
	return vectors
}

type VectorWithDistance struct {
	position utils.VectorI
	distance int
}

func closestUnvisited(nodeHeap *utils.Heap[VectorWithDistance], visited utils.Set[utils.VectorI]) utils.VectorI {
	for {
		if nodeHeap.Len() == 0 {
			return utils.VectorI{Down: -1}
		}
		next := nodeHeap.Pop()
		if !visited.Contains(next.position) {
			return next.position
		}
	}
}

func getNeighbours(v utils.VectorI, blockers utils.Set[utils.VectorI]) []utils.VectorI {
	neighbours := make([]utils.VectorI, 0, 4)
	for _, n := range v.GetCardinalAdjacents() {
		if utils.WithinGridSize(n, gridSize) && !blockers.Contains(n) {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func dijkstra(start utils.VectorI, end utils.VectorI, blocks utils.Set[utils.VectorI]) (int, bool) {
	visited := utils.NewSet[utils.VectorI]()
	distances := make(map[utils.VectorI]int)
	unvisited := utils.NewHeap[VectorWithDistance](func(a, b VectorWithDistance) bool { return a.distance < b.distance })
	distances[start] = 0

	currentNode := start
	for {
		neighbours := getNeighbours(currentNode, blocks)
		for _, n := range neighbours {
			if visited.Contains(n) {
				continue
			}
			distance := distances[currentNode] + 1
			bestDistance, ok := distances[n]
			if !ok || distance <= bestDistance {
				distances[n] = distance
				unvisited.Push(VectorWithDistance{position: n, distance: distance})
			}
		}
		visited.Add(currentNode)
		if currentNode == end {
			break
		}
		nextNode := closestUnvisited(unvisited, visited)
		if nextNode.Down == -1 {
			return -1, false
		}
		currentNode = nextNode
	}
	//grid := utils.MakeGrid(gridSize)
	//for v, distance := range distances {
	//	grid.Set(v, uint8(distance))
	//}
	//fmt.Println(&grid)
	return distances[end], true
}
