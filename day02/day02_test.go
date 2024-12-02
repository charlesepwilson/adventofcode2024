package day02

import (
	"testing"
)

func TestLineIsSafe(t *testing.T) {
	safeExamples := [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 6, 7, 9},
	}
	unSafeExamples := [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
	}

	for _, example := range safeExamples {
		if !lineIsSafe(example) {
			t.Errorf("Line %d is not safe", example)
		}
	}
	for _, example := range unSafeExamples {
		if lineIsSafe(example) {
			t.Errorf("Line %d is safe", example)
		}
	}
}

func TestDampenedLineIsSafe(t *testing.T) {
	safeExamples := [][]int{
		{7, 6, 4, 2, 1},
		{1, 3, 6, 7, 9},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
	}
	unSafeExamples := [][]int{
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
	}

	for _, example := range safeExamples {
		if !dampenedLineIsSafe(example) {
			t.Errorf("Line %d is not safe", example)
		}
	}
	for _, example := range unSafeExamples {
		if dampenedLineIsSafe(example) {
			t.Errorf("Line %d is safe", example)
		}
	}
}
