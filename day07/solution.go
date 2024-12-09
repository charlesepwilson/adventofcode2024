package day07

import (
	"advent_of_code_2024/utils"
	"bytes"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 7 }

func (Solution) Part1(input []byte) int {
	equations := parseInput(input)
	functionOptions := []BinaryIntFunc{Add, Mul}
	return sumViableEquations(equations, functionOptions)
}

func (Solution) Part2(input []byte) int {
	equations := parseInput(input)
	functionOptions := []BinaryIntFunc{Add, Mul, Concat}
	return sumViableEquations(equations, functionOptions)
}

func (Solution) GetExample() []byte {
	return []byte("190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20")
}

func (Solution) ExampleAnswer1() int {
	return 3749
}
func (Solution) ExampleAnswer2() int {
	return 11387
}

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

func countDigits(i int) int {
	count := 0
	for {
		if i == 0 {
			return count
		}
		i /= 10
		count += 1

	}
}

func Concat(a int, b int) int {
	// 111 || 222 = 111222 = 111 * 1000 + 222
	digits := countDigits(b)
	shiftMul := utils.Pow(10, digits)
	return (a * shiftMul) + b
}

func intToSliceIndexed2[T any](integer int, length int, options []T) []T {
	result := make([]T, length)
	for i := range result {
		bit := (integer >> i) & 1
		result[i] = options[bit]
	}
	return result
}

func convertBase(value int, base int, length int) []int {
	u := value
	a := make([]int, length)
	i := -1
	b := base
	//fmt.Println(u, base, i, a)
	for u >= b {
		//fmt.Println(u, base, i, a)
		i++
		// Avoid using r = a%b in addition to q = a/b
		// since 64bit division and modulo operations
		// are calculated by runtime functions on 32bit machines.
		q := u / b
		a[i] = u - q*b
		u = q
	}
	// u < base
	i++
	a[i] = u
	//fmt.Println(a)
	return a
}

func intToSliceIndexed[T any](integer int, length int, options []T) []T {
	base := len(options)
	if base == 2 {
		return intToSliceIndexed2(integer, length, options)
	}
	convertedBase := convertBase(integer, base, length)
	result := make([]T, length)
	for i := range result {
		index := convertedBase[i]
		result[i] = options[index]
	}
	return result
}

func (e Equation) couldWork(functionOptions []BinaryIntFunc) bool {
	operatorGaps := len(e.arguments) - 1
	numPossibilities := utils.Pow(len(functionOptions), operatorGaps) //  2 << operatorGaps
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

func sumViableEquations(equations []Equation, functionOptions []BinaryIntFunc) int {
	total := 0
	for _, equation := range equations {
		if equation.couldWork(functionOptions) {
			total += equation.answer
		}
	}
	return total
}
