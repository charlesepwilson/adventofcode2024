package day15

import (
	"advent_of_code_2024/utils"
	"bytes"
	"strings"
)

type Solution struct{}

func (Solution) Day() int { return 15 }

func (Solution) Part1(input []byte) int {
	grid, instructions, robotPos := parseInput(input)
	for _, instruction := range instructions {
		robotPos, grid = followInstruction(robotPos, grid, instruction)
	}
	result := 0
	for i, line := range grid {
		for j, item := range line {
			if item == BOX {
				result += 100*i + j
			}
		}
	}
	return result
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample(part int) []byte {
	return []byte("##########\n#..O..O.O#\n#......O.#\n#.OO..O.O#\n#..O@..O.#\n#O#..O...#\n#O..O..O.#\n#.OO.O.OO#\n#....O...#\n##########\n\n<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^\nvvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v\n><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<\n<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^\n^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><\n^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^\n>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^\n<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>\n^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>\nv^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^")
}

func (Solution) ExampleAnswer1() int {
	return 10092
}
func (Solution) ExampleAnswer2() int {
	return 0
}

const ROBOT = '@'
const BOX = 'O'
const WALL = '#'
const OUTSIDE = ';'

func parseInput(input []byte) (grid Grid, instructions []utils.VectorI, start utils.VectorI) {
	gridPart, instructionsPart, _ := bytes.Cut(input, []byte("\n\n"))

	grid = bytes.Split(gridPart, []byte("\n"))

	instructions = make([]utils.VectorI, 0, len(instructionsPart))
	for _, instruction := range instructionsPart {
		if instruction == '<' {
			instructions = append(instructions, utils.VectorI{Right: -1})
		} else if instruction == '>' {
			instructions = append(instructions, utils.VectorI{Right: 1})
		} else if instruction == 'v' {
			instructions = append(instructions, utils.VectorI{Down: 1})
		} else if instruction == '^' {
			instructions = append(instructions, utils.VectorI{Down: -1})
		}
	}

	for i, line := range grid {
		for j, obj := range line {
			if obj == ROBOT {
				start = utils.VectorI{Down: i, Right: j}
			}
		}
	}

	return grid, instructions, start
}

type Grid [][]byte

func (g *Grid) Get(v utils.VectorI) byte {
	if !utils.WithinGrid(v, g.Size()) {
		return OUTSIDE
	}
	return (*g)[v.Down][v.Right]
}

func (g *Grid) Set(v utils.VectorI, c byte) {
	if !utils.WithinGrid(v, g.Size()) {
		return
	}
	(*g)[v.Down][v.Right] = c
}

func (g *Grid) Size() utils.VectorI {
	return utils.VectorI{
		Down:  len(*g),
		Right: len((*g)[0]),
	}
}

func (g *Grid) String() string {
	builder := strings.Builder{}
	for _, row := range *g {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	return builder.String()
}

func followInstruction(robotPos utils.VectorI, grid Grid, instruction utils.VectorI) (newRobotPos utils.VectorI, newGrid Grid) {
	newRobotPos = robotPos.Add(instruction)
	itemAtNewRobotPos := grid.Get(newRobotPos)
	if itemAtNewRobotPos == WALL {
		return robotPos, grid
	} else if itemAtNewRobotPos == BOX {
		newBoxPos := newRobotPos
		for {
			newBoxPos = newBoxPos.Add(instruction)
			nextItem := grid.Get(newBoxPos)
			if nextItem == WALL {
				return robotPos, grid
			}
			if nextItem != BOX {
				grid.Set(newRobotPos, ROBOT)
				grid.Set(newBoxPos, BOX)
				return newRobotPos, grid
			}
		}
	} else {
		grid.Set(newRobotPos, ROBOT)
		return newRobotPos, grid
	}
}
