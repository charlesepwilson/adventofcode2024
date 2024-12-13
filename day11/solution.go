package day11

import (
	"bytes"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 11 }

func (Solution) Part1(input []byte) int {
	stones := inputToStones(input)
	return countAfter(stones, 25)
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample() []byte {
	return []byte("125 17")
}

func (Solution) ExampleAnswer1() int {
	return 55312
}
func (Solution) ExampleAnswer2() int {
	return 0
}

type Stone struct {
	Str []byte
	Int int
}

func stoneFromStr(str []byte) Stone {
	i, _ := strconv.Atoi(string(str))
	str = []byte(strconv.Itoa(i))
	return Stone{
		Str: str,
		Int: i,
	}
}

func stoneFromInt(i int) Stone {
	return Stone{
		Str: []byte(strconv.Itoa(i)),
		Int: i,
	}
}

var NullStone = Stone{[]byte("-1"), -1}

func (s Stone) Blink() (Stone, Stone) {
	if s.Int == 0 {
		return Stone{Str: []byte("1"), Int: 1}, NullStone
	} else if len(s.Str)%2 == 0 {
		midpoint := len(s.Str) / 2
		return stoneFromStr(s.Str[:midpoint]),
			stoneFromStr(s.Str[midpoint:])
	} else {
		return stoneFromInt(s.Int * 2024), NullStone
	}
}

func (s Stone) String() string {
	return string(s.Str)
}

func inputToStones(input []byte) []Stone {
	tokens := bytes.Fields(input)
	stones := make([]Stone, len(tokens))
	for i, token := range tokens {
		stones[i] = stoneFromStr(token)
	}
	return stones
}

func blinkStones(stones []Stone) []Stone {
	newStones := make([]Stone, 0, len(stones))
	for _, stone := range stones {
		left, right := stone.Blink()
		newStones = append(newStones, left)
		if right.Int != NullStone.Int {
			newStones = append(newStones, right)
		}
	}
	return newStones
}

func countAfter(stones []Stone, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = blinkStones(stones)
	}
	return len(stones)
}
