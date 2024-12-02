package utils

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func ReadInput(day int, part int) []byte {
	fileName := fmt.Sprintf("inputs/day%02dp%d.txt", day, part)
	fileData, err := os.ReadFile(fileName)
	if err == nil {
		return fileData
	}
	panic("Could not read file: " + fileName)
}

func GetTokenLines(day int, part int) [][][]byte {
	input := ReadInput(day, part)
	lines := bytes.Split(input, []byte("\n"))
	tokenLines := make([][][]byte, len(lines))

	for i, line := range lines {
		tokenLines[i] = bytes.Fields(line)
	}
	return tokenLines
}

func GetIntegerTokenLines(day int, part int) [][]int {
	tokenLines := GetTokenLines(day, part)
	integerLines := make([][]int, len(tokenLines))
	for i, line := range tokenLines {
		integerLines[i] = make([]int, len(line))
		for j, v := range line {
			integerLines[i][j], _ = strconv.Atoi(string(v))
		}
	}
	return integerLines
}
