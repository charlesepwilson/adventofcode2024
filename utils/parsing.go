package utils

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func ReadInput(day int, part int) []byte {
	fileName := fmt.Sprintf("inputs/day%02dp%d.txt", day, part)
	fileData, err := os.ReadFile(fileName)
	if err == nil {
		return fileData
	}
	panic("Could not read file: " + fileName)
}

func GetTokenLines(input []byte) [][][]byte {
	lines := bytes.Split(input, []byte("\n"))
	tokenLines := make([][][]byte, len(lines))

	for i, line := range lines {
		tokenLines[i] = bytes.Fields(line)
	}
	return tokenLines
}

func GetIntegerTokenLines(input []byte) [][]int {
	tokenLines := GetTokenLines(input)
	integerLines := make([][]int, len(tokenLines))
	for i, line := range tokenLines {
		integerLines[i] = make([]int, len(line))
		for j, v := range line {
			integerLines[i][j], _ = strconv.Atoi(string(v))
		}
	}
	return integerLines
}

func GetIntegerGrid(input []byte) [][]int {
	lines := bytes.Split(input, []byte("\n"))
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, v := range line {
			grid[i][j], _ = strconv.Atoi(string(v))
		}
	}
	return grid
}

type DaySolution interface {
	Part1(input []byte) int
	Part2(input []byte) int
	Day() int
	GetExample(part int) []byte
	ExampleAnswer1() int
	ExampleAnswer2() int
}

const fmtString = "day %02d part %d: %d\n"

func Part1(day DaySolution) int {
	input := ReadInput(day.Day(), 1)
	return day.Part1(input)
}
func Part2(day DaySolution) int {
	input := ReadInput(day.Day(), 1)
	return day.Part2(input)
}

func PrintPart1(day DaySolution) {
	fmt.Printf(fmtString, day.Day(), 1, Part1(day))
}
func PrintPart2(day DaySolution) {
	fmt.Printf(fmtString, day.Day(), 2, Part2(day))
}
func Print(day DaySolution) {
	PrintPart1(day)
	PrintPart2(day)
}

func TestPart1(day DaySolution, t *testing.T) {
	exampleInput := day.GetExample(1)
	result := day.Part1(exampleInput)
	if result != day.ExampleAnswer1() {
		t.Errorf("Wrong answer for day %d p1: %d", day.Day(), result)
	}
}

func TestPart2(day DaySolution, t *testing.T) {
	exampleInput := day.GetExample(2)
	result := day.Part2(exampleInput)
	if result != day.ExampleAnswer2() {
		t.Errorf("Wrong answer for day %d p2: %d", day.Day(), result)
	}
}
