package day04

import (
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	example := []byte("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
	lines := bytes.Split(example, []byte("\n"))
	result := countXmas(lines)
	if result != 18 {
		t.Errorf("Wrong answer for day 4 p1: %d", result)
	}
}

func TestPart2(t *testing.T) {
	example := []byte("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
	lines := bytes.Split(example, []byte("\n"))
	result := countXshapedMas(lines)
	if result != 9 {
		t.Errorf("Wrong answer for day 4 p2: %d", result)
	}
}
