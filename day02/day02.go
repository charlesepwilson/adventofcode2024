package day02

import (
	"advent_of_code_2024/utils"
)

const DAY = 2

const MaxDiff = 3

func eLineIsSafe(line []int, maxErrors int) bool {
	errors := 0
	grad := 0
	for i := 1; i < len(line); i++ {
		g := line[i] - line[i-1]
		if g != 0 {
			grad += g / utils.Abs(g)
		}
	}
	if grad == 0 {
		return false
	}
	gradSign := grad / utils.Abs(grad)
	for i := 1; i < len(line); i++ {
		modDiff := (line[i] - line[i-1]) * gradSign
		if modDiff > MaxDiff || modDiff <= 0 {
			errors += 1
			if errors > maxErrors {
				return false
			} else {
				// whenever an error is found, there are 2 candidates for removal: left or right
				// I just make a new list and recursively call this function with one less error allowed
				// removing each of the candidates, but it feels like there should be a better way
				start := line[:i-1]
				end := line[i+1:]

				// I haven't got the hang of slice manipulation in Go yet; this 100% feels like there's a better way
				leftLine := make([]int, 0)
				leftLine = append(leftLine, start...)
				leftLine = append(leftLine, line[i-1])
				leftLine = append(leftLine, end...)

				rightLine := make([]int, 0)
				rightLine = append(rightLine, start...)
				rightLine = append(rightLine, line[i])
				rightLine = append(rightLine, end...)
				tryLeft := eLineIsSafe(leftLine, maxErrors-1)
				tryRight := eLineIsSafe(rightLine, maxErrors-1)
				return tryLeft || tryRight
			}
		}

	}
	return true
}

func lineIsSafe(line []int) bool {
	return eLineIsSafe(line, 0)
}

func Part1() int {
	integerLines := utils.GetIntegerTokenLines(DAY, 1)
	numSafe := 0
	for _, line := range integerLines {
		if lineIsSafe(line) {
			numSafe += 1
		}
	}
	return numSafe
}

func dampenedLineIsSafe(line []int) bool {
	return eLineIsSafe(line, 1)
}

func Part2() int {
	integerLines := utils.GetIntegerTokenLines(DAY, 1)
	numSafe := 0
	for _, line := range integerLines {
		if dampenedLineIsSafe(line) {
			numSafe += 1
		}
	}
	return numSafe
}
