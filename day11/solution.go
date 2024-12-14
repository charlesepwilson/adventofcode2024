package day11

import (
	"bytes"
	"fmt"
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 11 }

func (Solution) Part1(input []byte) int {
	stones := inputToStones(input)
	return countAfter(stones, 25)
}

func (Solution) Part2(input []byte) int {
	stones := inputToStones(input)
	return countAfter(stones, 75)
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

type Stone int

func (s Stone) Str() []byte {
	return []byte(strconv.Itoa(int(s)))
}

func stoneFromStr(str []byte) Stone {
	i, _ := strconv.Atoi(string(str))
	return Stone(i)
}

func stoneFromInt(i int) Stone {
	return Stone(i)
}

var NullStone = Stone(-1)

func (s Stone) Blink() (Stone, Stone) {
	if s == 0 {
		return Stone(1), NullStone
	} else if str := s.Str(); len(str)%2 == 0 {
		midpoint := len(str) / 2
		return stoneFromStr(str[:midpoint]),
			stoneFromStr(str[midpoint:])
	} else {
		return stoneFromInt(int(s) * 2024), NullStone
	}
}

func (s Stone) String() string {
	return string(s.Str())
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
		if right != NullStone {
			newStones = append(newStones, right)
		}
	}
	return newStones
}

func countAfter(stones []Stone, blinks int) int {
	total := 0
	for _, stone := range stones {
		total += len(blinkStone(stone, blinks))
	}
	return total
}

type BlinkCache map[int][]Stone

var bigBlinkCache = make(map[Stone]BlinkCache) // todo try cacheing but just counts rather than lists

func blinkStone(stone Stone, numTimes int) []Stone {
	fmt.Println(bigBlinkCache)
	if numTimes == 0 {
		return []Stone{stone}
	}
	result := make([]Stone, 0)
	blinkCache, ok := bigBlinkCache[stone]
	if !ok {
		blinkCache = make(BlinkCache)
		bigBlinkCache[stone] = blinkCache
	}
	defer func() { blinkCache[numTimes] = result }()
	for i := numTimes; i > 0; i-- {
		afterBlinks, okk := blinkCache[i]
		if okk {
			for _, s := range afterBlinks {
				result = append(result, blinkStone(s, numTimes-i)...)
			}
			blinkCache[i] = result
			return result
		}
	}
	left, right := stone.Blink()
	blinkCache[1] = []Stone{left}
	fullLeft := blinkStone(left, numTimes-1)
	result = append(result, fullLeft...)

	if right != NullStone {
		blinkCache[1] = append(blinkCache[1], right)
		fullRight := blinkStone(right, numTimes-1)
		result = append(result, fullRight...)
	}
	return result
}
