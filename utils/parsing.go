package utils

import (
	"fmt"
	"os"
)

func ReadInput(day int, part int) []byte {
	fileName := fmt.Sprintf("inputs/day%02dp%d.txt", day, part)
	fileData, err := os.ReadFile(fileName)
	if err == nil {
		return fileData
	}
	panic("Could not read file: " + fileName)
}
