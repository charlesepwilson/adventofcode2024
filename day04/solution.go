package day04

import (
	"bytes"
	"sort"
)

type Solution struct{}

func (Solution) Day() int { return 4 }

func (Solution) Part1(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	return countXmas(lines)
}

func (Solution) Part2(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	return countXshapedMas(lines)
}

func (Solution) GetExample() []byte {
	return []byte("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
}

func (Solution) ExampleAnswer1() int {
	return 18
}
func (Solution) ExampleAnswer2() int {
	return 9
}

func isMatch(input [][]byte, startRow int, startCol int, right int, down int, match []byte) bool {
	for i, letter := range match {
		nextRow := startRow + i*down
		nextCol := startCol + i*right
		if nextRow < 0 || nextRow >= len(input) || nextCol < 0 || nextCol >= len(input[0]) {
			return false
		}
		if input[nextRow][nextCol] != letter {
			return false
		}
	}
	return true

}

func countXmas(lines [][]byte) int {
	match := []byte("XMAS")
	total := 0
	for i, row := range lines {
		for j := range row {
			for y := -1; y <= 1; y += 1 {
				for x := -1; x <= 1; x += 1 {
					if !(x == 0 && y == 0) && isMatch(lines, i, j, x, y, match) {
						total += 1
					}
				}
			}
		}
	}
	return total
}

func sliceIsMas(slice []byte) bool {
	if slice[1] != 'A' {
		return false
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return bytes.Equal(slice, []byte("AMS"))
}

func boxIsMas(box [][]byte) bool {
	neg := []byte{box[0][0], box[1][1], box[2][2]}
	pos := []byte{box[0][2], box[1][1], box[2][0]}
	return sliceIsMas(pos) && sliceIsMas(neg)
}

func countXshapedMas(lines [][]byte) int {
	total := 0
	for i, row := range lines {
		for j := range row {
			size := 3
			rightSide := j + size
			bottomSide := i + size

			if rightSide <= len(row) && bottomSide <= len(lines) {
				box := [][]byte{
					lines[i][j:rightSide],
					lines[i+1][j:rightSide],
					lines[i+2][j:rightSide],
				}
				if boxIsMas(box) {
					total += 1
				}
			}
		}

	}
	return total
}
