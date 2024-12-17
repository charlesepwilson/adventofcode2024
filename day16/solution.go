package day16

import (
	"advent_of_code_2024/utils"
	"bytes"
)

type Solution struct{}

func (Solution) Day() int { return 16 }

func (Solution) Part1(input []byte) int {
	start, end, grid := parseInput(input)
	_, minDist := dijkstra(start, end, grid)
	return minDist
}

func (Solution) Part2(input []byte) int {
	//start, end, grid := parseInput(input)
	//distances, minDist := dijkstra(start, end, grid)

	return len(input)
}

func (Solution) GetExample(part int) []byte {
	return []byte("###############\n#.......#....E#\n#.#.###.#.###.#\n#.....#.#...#.#\n#.###.#####.#.#\n#.#.#.......#.#\n#.#.#####.###.#\n#...........#.#\n###.#.#####.#.#\n#...#.....#.#.#\n#.#.#.###.#.#.#\n#.....#...#.#.#\n#.###.#.#.#.#.#\n#S..#.....#...#\n###############")
}

func (Solution) ExampleAnswer1() int {
	return 7036
}
func (Solution) ExampleAnswer2() int {
	return 45
}

func parseInput(input []byte) (start Node, end utils.VectorI, grid utils.Grid) {
	grid = bytes.Split(input, []byte("\n"))
	for i, line := range grid {
		for j, char := range line {
			if char == 'S' {
				start = Node{
					position:  utils.VectorI{Down: i, Right: j},
					direction: 1,
				}
			} else if char == 'E' {
				end = utils.VectorI{Down: i, Right: j}
			}
		}
	}
	return
}

type Node struct {
	position  utils.VectorI
	direction int
}

const costTurn = 1000
const costMove = 1
const WALL = '#'

func canMoveTo(position utils.VectorI, grid utils.Grid) bool {
	itemAtMovePos := grid.Get(position)
	return itemAtMovePos != WALL && itemAtMovePos != utils.OUTSIDE
}

func (n Node) getNeighbours(grid utils.Grid) map[Node]int {
	neighbours := make(map[Node]int, 3)
	straight := Node{
		position:  n.position.Add(utils.Directions[n.direction]),
		direction: n.direction,
	}
	if canMoveTo(straight.position, grid) {
		neighbours[straight] = costMove
	}
	numDirs := len(utils.Directions)
	right := Node{
		position:  n.position,
		direction: (n.direction + 1) % numDirs,
	}
	neighbours[right] = costTurn
	left := Node{
		position:  n.position,
		direction: (n.direction + numDirs - 1) % numDirs,
	}
	neighbours[left] = costTurn

	return neighbours
}

func closestUnvisited(nodeHeap *utils.Heap[NodeWithDistance], visited utils.Set[Node]) Node {
	for {
		next := nodeHeap.Pop()
		if !visited.Contains(next.node) {
			return next.node
		}
	}
}

type NodeWithDistance struct {
	node Node
	dist int
}

func dijkstra(start Node, end utils.VectorI, grid utils.Grid) (map[Node]int, int) {
	visited := utils.NewSet[Node]()
	distances := make(map[Node]int)
	unvisited := utils.NewHeap[NodeWithDistance](func(a, b NodeWithDistance) bool { return a.dist < b.dist })
	distances[start] = 0
	currentNode := start
	endNodes := []Node{
		{direction: 0, position: end},
		{direction: 1, position: end},
	}
	for {
		finished := true
		for _, endNode := range endNodes {
			if !visited.Contains(endNode) {
				finished = false
				break
			}
		}
		if finished {
			break
		}

		neighbours := currentNode.getNeighbours(grid)
		for n, d := range neighbours {
			if visited.Contains(n) {
				continue
			}
			distance := distances[currentNode] + d
			bestDistance, ok := distances[n]
			if !ok || distance < bestDistance {
				distances[n] = distance
				unvisited.Push(NodeWithDistance{n, distance})
			}
		}
		visited.Add(currentNode)
		grid.Set(currentNode.position, 'o')
		nextNode := closestUnvisited(unvisited, visited)
		currentNode = nextNode
	}

	return distances, utils.Min(distances[endNodes[0]], distances[endNodes[1]])
}
