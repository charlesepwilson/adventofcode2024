package day14

import (
	"advent_of_code_2024/utils"
	"bytes"
	"fmt"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 14 }

func (Solution) Part1(input []byte) int {
	robots := parseInput(input)
	for i := range robots {
		robots[i] = robots[i].Advance(100)
	}
	robotsPerQuadrant := countRobotsPerQuadrant(robots)
	result := 1
	for _, q := range robotsPerQuadrant {
		result *= q
	}
	return result
}

var iteration = 0

func (Solution) Part2(input []byte) int {
	robots := parseInput(input)
	advRate := 1
	for {
		for i := range robots {
			robots[i] = robots[i].Advance(advRate)
		}
		iteration += advRate
		if looksChristmassy(robots) {
			return iteration
		}
	}

}

func (Solution) GetExample(part int) []byte {
	gridSize = utils.VectorI{
		Right: 11,
		Down:  7,
	}
	return []byte("p=0,4 v=3,-3\np=6,3 v=-1,-3\np=10,3 v=-1,2\np=2,0 v=2,-1\np=0,0 v=1,3\np=3,0 v=-2,-2\np=7,6 v=-1,-3\np=3,0 v=-1,-2\np=9,3 v=2,3\np=7,3 v=-1,2\np=2,4 v=2,-3\np=9,5 v=-3,-3")
}

func (Solution) ExampleAnswer1() int {
	return 12
}
func (Solution) ExampleAnswer2() int {
	return 0
}

var gridSize = utils.VectorI{
	Right: 101,
	Down:  103,
}

type Robot struct {
	position, velocity utils.VectorI
}

func (r Robot) Advance(seconds int) Robot {
	newPosition := r.position.Add(r.velocity.Mul(seconds))
	newPosition.Right = utils.EuclideanMod(newPosition.Right, gridSize.Right)
	newPosition.Down = utils.EuclideanMod(newPosition.Down, gridSize.Down)
	r.position = newPosition
	return Robot{position: newPosition, velocity: r.velocity}
}

func parseInput(input []byte) []Robot {
	lines := bytes.Split(input, []byte("\n"))
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		p, v, _ := bytes.Cut(line, []byte(" "))
		pxs, pys, _ := bytes.Cut(p[2:], []byte(","))
		vxs, vys, _ := bytes.Cut(v[2:], []byte(","))
		px, _ := strconv.Atoi(string(pxs))
		py, _ := strconv.Atoi(string(pys))
		vx, _ := strconv.Atoi(string(vxs))
		vy, _ := strconv.Atoi(string(vys))
		robots[i].position.Right = px
		robots[i].position.Down = py
		robots[i].velocity.Right = vx
		robots[i].velocity.Down = vy
	}
	return robots
}

func getQuadrant(v utils.VectorI) (int, bool) {
	var horizontalMidpoint = gridSize.Right / 2
	var verticalMidpoint = gridSize.Down / 2
	if v.Right == horizontalMidpoint || v.Down == verticalMidpoint {
		return -1, false
	}
	q := 0
	if v.Right > horizontalMidpoint {
		q++
	}
	if v.Down > verticalMidpoint {
		q += 2
	}
	return q, true
}

func countRobotsPerQuadrant(robots []Robot) [4]int {
	counts := [4]int{}
	for _, robot := range robots {
		quadrant, valid := getQuadrant(robot.position)
		if valid {
			counts[quadrant]++
		}
	}
	return counts
}

func next(b byte) byte {
	if b == ' ' {
		return '#'
	} else if b == '#' {
		return '*'
	} else {
		return '+'
	}
}

func printRobots(robots []Robot) {
	fmt.Println("Iteration", iteration)
	grid := make([][]byte, gridSize.Down)
	for i := range grid {
		grid[i] = make([]byte, gridSize.Right)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	for _, robot := range robots {
		grid[robot.position.Down][robot.position.Right] = next(grid[robot.position.Down][robot.position.Right])
	}
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func looksChristmassy(robots []Robot) bool {
	var horizontalMidpoint = gridSize.Right / 2
	gradient := 2
	numInTriangle := 0
	for _, robot := range robots {
		var h int
		if robot.position.Right < horizontalMidpoint {
			h = robot.position.Right
		} else {
			h = gridSize.Right - robot.position.Right
		}
		if robot.position.Down > gridSize.Down-(gradient*h) {
			numInTriangle++
		}
	}
	if (1000*numInTriangle)/len(robots) > 750 {
		printRobots(robots)
		return true
	} else {
		return false
	}
}
