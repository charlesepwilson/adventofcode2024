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
	return computeTotalGps(grid, BOX)
}

func (Solution) Part2(input []byte) int {
	grid, instructions, robotPos := parseInput(input)
	grid = widenGrid(grid)
	robotPos.Right = 2 * robotPos.Right
	for _, instruction := range instructions {
		//fmt.Println("robot", robotPos, "dir", instruction)
		//fmt.Println(&grid)
		robotPos, grid = followWideInstruction(robotPos, grid, instruction)
	}
	return computeTotalGps(grid, LEFTBOX)
}

func (Solution) GetExample(part int) []byte {
	return []byte("##########\n#..O..O.O#\n#......O.#\n#.OO..O.O#\n#..O@..O.#\n#O#..O...#\n#O..O..O.#\n#.OO.O.OO#\n#....O...#\n##########\n\n<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^\nvvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v\n><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<\n<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^\n^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><\n^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^\n>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^\n<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>\n^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>\nv^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^")
}

func (Solution) ExampleAnswer1() int {
	return 10092
}
func (Solution) ExampleAnswer2() int {
	return 9021
}

const ROBOT = '@'
const BOX = 'O'
const WALL = '#'
const OUTSIDE = ';'
const LEFTBOX = '['
const RIGHTBOX = ']'

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

func widenGrid(grid Grid) Grid {
	newGrid := make(Grid, len(grid))
	for i, row := range grid {
		newRow := make([]byte, 0, len(row)*2)
		for _, char := range row {
			var n []byte
			if char == BOX {
				n = []byte{LEFTBOX, RIGHTBOX}
			} else {
				n = []byte{char, char}
			}
			newRow = append(newRow, n...)
		}
		newGrid[i] = newRow
	}
	return newGrid
}

func followWideInstruction(robotPos utils.VectorI, grid Grid, instruction utils.VectorI) (newRobotPos utils.VectorI, newGrid Grid) {
	newRobotPos = robotPos.Add(instruction)
	itemAtNewRobotPos := grid.Get(newRobotPos)
	if itemAtNewRobotPos == WALL {
		return robotPos, grid
	} else if itemAtNewRobotPos == LEFTBOX || itemAtNewRobotPos == RIGHTBOX {
		success := canPush(newRobotPos, grid, instruction)
		if !success {
			return robotPos, grid
		} else {
			applyPush(newRobotPos, grid, instruction)
			grid.Set(newRobotPos, ROBOT)
			return newRobotPos, grid
		}
	} else {
		grid.Set(newRobotPos, ROBOT)
		return newRobotPos, grid
	}
}

func canPush(pushObjAt utils.VectorI, grid Grid, instruction utils.VectorI) bool {
	obj := grid.Get(pushObjAt)
	if obj == WALL {
		return false
	} else if obj == LEFTBOX || obj == RIGHTBOX {
		pushInto := pushObjAt.Add(instruction)
		if instruction.Right == 0 {
			// vertical; need to handle wide boxes
			var offset int
			if obj == LEFTBOX {
				offset = 1
			} else {
				offset = -1
			}
			otherBoxPart := pushObjAt.Add(utils.VectorI{Right: offset})
			otherPushInto := otherBoxPart.Add(instruction)
			success := canPush(pushInto, grid, instruction) && canPush(otherPushInto, grid, instruction)
			return success
		} else {
			return canPush(pushInto, grid, instruction)
		}
	} else {
		return true
	}
}

func applyPush(pushObjAt utils.VectorI, grid Grid, instruction utils.VectorI) {
	obj := grid.Get(pushObjAt)
	if obj == LEFTBOX || obj == RIGHTBOX {
		pushInto := pushObjAt.Add(instruction)
		applyPush(pushInto, grid, instruction)
		grid.Set(pushInto, obj)
		grid.Set(pushObjAt, ROBOT)
		if instruction.Right == 0 {
			// vertical; need to handle wide boxes
			var offset int
			if obj == LEFTBOX {
				offset = 1
			} else {
				offset = -1
			}
			otherBoxPart := pushObjAt.Add(utils.VectorI{Right: offset})
			otherPushInto := otherBoxPart.Add(instruction)
			applyPush(otherPushInto, grid, instruction)
			grid.Set(otherPushInto, grid.Get(otherBoxPart))
			grid.Set(otherBoxPart, '.')
		}
	}
}

func computeTotalGps(grid Grid, target byte) int {
	result := 0
	for i, line := range grid {
		for j, item := range line {
			if item == target {
				result += 100*i + j
			}
		}
	}
	return result
}
