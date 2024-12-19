package day19

import (
	"bytes"
	"slices"
)

type Solution struct{}

func (Solution) Day() int { return 19 }

func (Solution) Part1(input []byte) int {
	patterns, designs := parseInput(input)
	slices.SortFunc[[][]byte](patterns, func(a, b []byte) int {
		if len(a) == len(b) {
			return 0
		} else if len(a) < len(b) {
			return -1
		} else {
			return 1
		}
	})
	patterns = simplifyPatterns(patterns)
	//printPatterns(patterns)
	result := 0
	for _, design := range designs {
		//fmt.Println(string(design))
		if isPossibleDesign(patterns, design) {
			result++
		}
	}
	return result
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample(part int) []byte {
	return []byte("r, wr, b, g, bwu, rb, gb, br\n\nbrwrr\nbggr\ngbbr\nrrbgbr\nubwu\nbwurrg\nbrgr\nbbrgwb")
}

func (Solution) ExampleAnswer1() int {
	return 6
}
func (Solution) ExampleAnswer2() int {
	return 0
}

func parseInput(input []byte) (patterns [][]byte, designs [][]byte) {
	patternPart, designPart, _ := bytes.Cut(input, []byte("\n\n"))
	patterns = bytes.Split(patternPart, []byte(", "))
	designs = bytes.Split(designPart, []byte("\n"))
	return patterns, designs
}

func isPossibleDesign(patterns [][]byte, design []byte) bool {
	if len(design) == 0 {
		return true
	}
	for _, pattern := range patterns {
		if len(design) >= len(pattern) && slices.Equal(design[:len(pattern)], pattern) {
			if isPossibleDesign(patterns, design[len(pattern):]) {
				return true
			}
		}
	}
	return false
}

//func printPatterns(patterns [][]byte) {
//	ps := make([]string, len(patterns))
//	for i, pattern := range patterns {
//		ps[i] = string(pattern)
//	}
//	fmt.Println(strings.Join(ps, ", "))
//}

func simplifyPatterns(patterns [][]byte) [][]byte {
	//printPatterns(patterns)
	for i := len(patterns) - 1; i >= 0; i-- {
		patternsExcl := make([][]byte, i, len(patterns)-1)
		copy(patternsExcl, patterns[:i])
		if i < len(patterns)-1 {
			patternsExcl = append(patternsExcl, patterns[i+1:]...)
		}
		//printPatterns(patternsExcl)
		//fmt.Println(string(patterns[i]))
		//fmt.Println()
		if isPossibleDesign(patternsExcl, patterns[i]) {
			patterns = patternsExcl
		}
	}
	return patterns
}
