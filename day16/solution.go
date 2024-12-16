package day16

import (
	"advent_of_code_2024/utils"
	"bytes"
	"math"
)

type Solution struct{}

func (Solution) Day() int { return 16 }

func (Solution) Part1(input []byte) int {
	start, end, grid := parseInput(input)
	return dijkstra(start, end, grid)
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample(part int) []byte {
	return []byte("###############\n#.......#....E#\n#.#.###.#.###.#\n#.....#.#...#.#\n#.###.#####.#.#\n#.#.#.......#.#\n#.#.#####.###.#\n#...........#.#\n###.#.#####.#.#\n#...#.....#.#.#\n#.#.#.###.#.#.#\n#.....#...#.#.#\n#.###.#.#.#.#.#\n#S..#.....#...#\n###############")
}

func (Solution) ExampleAnswer1() int {
	return 7036
}
func (Solution) ExampleAnswer2() int {
	return 0
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

func (n Node) getNeighbours(grid utils.Grid) map[Node]int {
	neighbours := make(map[Node]int, 3)
	movePos := n.position.Add(utils.Directions[n.direction])
	itemAtMovePos := grid.Get(movePos)
	if itemAtMovePos != WALL && itemAtMovePos != utils.OUTSIDE {
		neighbours[Node{
			position:  movePos,
			direction: n.direction,
		}] = costMove
	}
	numDirs := len(utils.Directions)
	neighbours[Node{
		position:  n.position,
		direction: (n.direction + 1) % numDirs,
	}] = costTurn
	neighbours[Node{
		position:  n.position,
		direction: (n.direction + numDirs - 1) % numDirs,
	}] = costTurn
	return neighbours
}

func closestUnvisited(distances map[Node]int, visited utils.Set[Node]) Node {
	minDistance := math.MaxInt
	node := Node{}
	for n, d := range distances {
		if visited.Contains(n) {
			continue
		}
		if d < minDistance {
			minDistance = d
			node = n
		}
	}
	return node
}

func dijkstra(start Node, end utils.VectorI, grid utils.Grid) int {
	visited := utils.NewSet[Node]()
	distances := make(map[Node]int)
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
			}
		}
		visited.Add(currentNode)
		currentNode = closestUnvisited(distances, visited)
	}

	return utils.Min(distances[endNodes[0]], distances[endNodes[1]])
}
