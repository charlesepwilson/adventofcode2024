package day19

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

type Solution struct{}

func (Solution) Day() int { return 19 }

func (Solution) Part1(input []byte) int {
	patterns, designs := parseInput(input)
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
	patterns, designs := parseInput(input)
	patterns = sortPatterns(patterns)
	simplifiedPatterns := simplifyPatterns(patterns)
	//printPatterns(patterns)
	result := 0
	for _, design := range designs {
		//fmt.Println()
		//fmt.Println(string(design))
		if isPossibleDesign(simplifiedPatterns, design) {
			result += countWaysToBuild(design, patterns, 0)
		}

	}
	return result
}

func (Solution) GetExample(part int) []byte {
	return []byte("r, wr, b, g, bwu, rb, gb, br\n\nbrwrr\nbggr\ngbbr\nrrbgbr\nubwu\nbwurrg\nbrgr\nbbrgwb")
}

func (Solution) ExampleAnswer1() int {
	return 6
}
func (Solution) ExampleAnswer2() int {
	return 16
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

func printPatterns(patterns [][]byte) {
	ps := make([]string, len(patterns))
	for i, pattern := range patterns {
		ps[i] = string(pattern)
	}
	fmt.Println(strings.Join(ps, ", "))
}

func sortPatterns(patterns [][]byte) [][]byte {
	slices.SortFunc[[][]byte](patterns, func(a, b []byte) int {
		if len(a) == len(b) {
			return 0
		} else if len(a) > len(b) {
			return -1
		} else {
			return 1
		}
	})
	return patterns
}

func simplifyPatterns(patterns [][]byte) [][]byte {
	//patterns := make([][]byte, len(p))
	//copy(patterns, p)
	sortPatterns(patterns)
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

func indent(r int) {
	for i := 0; i < r; i++ {
		fmt.Print("  ")
	}
}

var countCache = make(map[string]int)

func countWaysToBuild(design []byte, patterns [][]byte, r int) int {
	if len(design) == 0 {
		return 1
	}
	if count, ok := countCache[string(design)]; ok {
		return count
	}
	total := 0
	for _, pattern := range patterns {
		//indent(r)
		//fmt.Println("design", string(design), "pattern", string(pattern))
		if len(design) >= len(pattern) && slices.Equal(design[:len(pattern)], pattern) {
			//indent(r)
			//fmt.Println("cont.")
			ways := countWaysToBuild(design[len(pattern):], patterns, r+1)

			//indent(r)
			//fmt.Println("ways", ways)
			total += ways
			//if isPossibleDesign(patterns, design[len(pattern):]) {
			//	total++
			//}
		}
	}
	countCache[string(design)] = total
	return total
}
