package day03

import (
	"testing"
)

func TestPart2(t *testing.T) {
	example := []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	result := totalWithDoDont(example)
	if result != 48 {
		t.Errorf("Wrong answer for day 3 p2: %d", result)
	}

}
