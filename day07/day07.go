package day07

import (
	"advent_of_code_2024/utils"
	"bytes"
	"strconv"
)

const DAY = 7

type Equation struct {
	answer    int
	arguments []int
}

func parseInput(input []byte) (equations []Equation) {
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		var equation Equation
		answerStr, argStr, _ := bytes.Cut(line, []byte(": "))
		equation.answer, _ = strconv.Atoi(string(answerStr))
		for _, arg := range bytes.Split(argStr, []byte(" ")) {
			num, _ := strconv.Atoi(string(arg))
			equation.arguments = append(equation.arguments, num)
		}
		equations = append(equations, equation)
	}
	return equations
}

type BinaryIntFunc func(a int, b int) int

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func intToSliceIndexed[T any](integer int, length int, options []T) []T {
	result := make([]T, length)
	for i := range result {
		bit := (integer >> i) & 1
		result[i] = options[bit]
	}
	return result
}

func (e Equation) couldWork() bool {
	functionOptions := []BinaryIntFunc{Add, Mul}
	operatorGaps := len(e.arguments) - 1
	numPossibilities := 2 << operatorGaps
	for i := 0; i < numPossibilities; i++ {
		functionList := intToSliceIndexed(i, operatorGaps, functionOptions)
		result := e.arguments[0]
		for gapIndex, function := range functionList {
			result = function(result, e.arguments[gapIndex+1])
		}
		if result == e.answer {
			return true
		}
	}
	return false
}

func part1(input []byte) int {
	equations := parseInput(input)
	total := 0
	for _, equation := range equations {
		if equation.couldWork() {
			total += equation.answer
		}
	}
	return total
}

func part2(input []byte) int {
	//startPos, obstacles, gridSize := parseInput(input)
	total := 0
	//fmt.Println(loopMakers)
	return total
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
