package day13

import (
	"advent_of_code_2024/utils"
	"bytes"
	"math"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 13 }

func (Solution) Part1(input []byte) int {
	machines := parseInput(input)
	return countTokensForAllWins(machines)
}

func (Solution) Part2(input []byte) int {
	machines := parseInput(input)
	valueIncrease := 10000000000000
	vectorIncrease := utils.VectorI{Down: valueIncrease, Right: valueIncrease}
	for i := range machines {
		machines[i].prize = machines[i].prize.Add(vectorIncrease)
	}
	return countTokensForAllWins(machines)
}

func (Solution) GetExample(part int) []byte {
	return []byte("Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279")
}

func (Solution) ExampleAnswer1() int {
	return 480
}
func (Solution) ExampleAnswer2() int {
	return 875318608908
}

const costA int = 3
const costB int = 1

type Button struct {
	utils.VectorI
}

func (b Button) X() int { return b.Right }
func (b Button) Y() int { return b.Down }

type Machine struct {
	buttonA Button
	buttonB Button
	prize   utils.VectorI
}

func parseButton(input []byte) Button {
	return Button{parsePrize(input)}
}

func parsePrize(input []byte) utils.VectorI {
	_, info, _ := bytes.Cut(input, []byte(": "))
	xPart, yPart, _ := bytes.Cut(info, []byte(", "))
	x, _ := strconv.Atoi(string(xPart[2:]))
	y, _ := strconv.Atoi(string(yPart[2:]))
	return utils.VectorI{Down: y, Right: x}
}

func parseInput(input []byte) []Machine {
	sections := bytes.Split(input, []byte("\n\n"))
	machines := make([]Machine, len(sections))
	for i, section := range sections {
		lines := bytes.Split(section, []byte("\n"))
		machines[i].buttonA = parseButton(lines[0])
		machines[i].buttonB = parseButton(lines[1])
		machines[i].prize = parsePrize(lines[2])
	}
	return machines
}

type Matrix2D struct {
	a, b, c, d float64
}

func (m Matrix2D) Inverse() Matrix2D {
	determinant := m.a*m.d - m.b*m.c
	return Matrix2D{
		m.d / determinant,
		-m.b / determinant,
		-m.c / determinant,
		m.a / determinant,
	}
}

type Vector2D struct {
	x, y float64
}

func (m Matrix2D) Multiply(v Vector2D) Vector2D {
	return Vector2D{
		x: m.a*v.x + m.b*v.y,
		y: m.c*v.x + m.d*v.y,
	}
}

func isPositiveInteger(f float64) bool {
	return utils.AlmostEqual(math.Round(f), f) && f >= 0
}

func getWinningButtons(machine Machine) (aCount, bCount int, success bool) {
	if aDir := machine.buttonA.Simplify(); aDir == machine.buttonB.Simplify() {
		panic("got parallel buttons")
		return aCount, bCount, false // todo complete this case
		//// parallel, so multiple potential solutions could exist
		//if aDir != machine.prize.Simplify() {
		//	return aCount, bCount, false
		//}
		//efficiencyA := float32(machine.buttonA.X()) / float32(costA)
		//efficiencyB := float32(machine.buttonB.X()) / float32(costB)
		//if efficiencyA > efficiencyB {
		//	// press A as far as possible then finish with b
		//} else {
		//	// press B as far as possible then finish with a
		//}
	}
	startVec := Vector2D{
		x: float64(machine.prize.Right),
		y: float64(machine.prize.Down),
	}
	basisChangeMatrix := Matrix2D{
		a: float64(machine.buttonA.X()),
		b: float64(machine.buttonB.X()),
		c: float64(machine.buttonA.Y()),
		d: float64(machine.buttonB.Y()),
	}.Inverse()
	finalVec := basisChangeMatrix.Multiply(startVec)
	if !isPositiveInteger(finalVec.x) || !isPositiveInteger(finalVec.y) {
		return aCount, bCount, false
	}

	return int(math.Round(finalVec.x)), int(math.Round(finalVec.y)), true
}

func getCost(aPresses int, bPresses int) int {
	return aPresses*costA + bPresses*costB
}

func countTokensForAllWins(machines []Machine) int {
	result := 0
	for _, machine := range machines {
		aPresses, bPresses, success := getWinningButtons(machine)
		if success {
			result += getCost(aPresses, bPresses)
		}
	}
	return result
}
