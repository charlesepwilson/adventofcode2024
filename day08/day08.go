package day08

import (
	"advent_of_code_2024/utils"
	"bytes"
	"strconv"
)

const DAY = 7

func parseInput(input []byte) {
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

func part1(input []byte) int {
	equations := parseInput(input)
	functionOptions := []BinaryIntFunc{Add, Mul}
	return sumViableEquations(equations, functionOptions)
}

func part2(input []byte) int {
	equations := parseInput(input)
	functionOptions := []BinaryIntFunc{Add, Mul, Concat}
	return sumViableEquations(equations, functionOptions)
}

func Part1() int {
	return part1(utils.ReadInput(DAY, 1))
}

func Part2() int {
	return part2(utils.ReadInput(DAY, 1))
}
