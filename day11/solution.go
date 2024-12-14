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
	return 65601038650482
}

type Stone int

func (s Stone) Str() []byte {
	return []byte(strconv.Itoa(int(s)))
}

func stoneFromStr(str []byte) Stone {
	i, _ := strconv.Atoi(string(str))
	return Stone(i)
}

var NullStone = Stone(-1)

func (s Stone) IsNull() bool {
	return s == NullStone
}

func (s Stone) Blink() (Stone, Stone) {
	if s == 0 {
		return Stone(1), NullStone
	} else if str := s.Str(); len(str)%2 == 0 {
		midpoint := len(str) / 2
		return stoneFromStr(str[:midpoint]),
			stoneFromStr(str[midpoint:])
	} else {
		return Stone(int(s) * 2024), NullStone
	}
}

func inputToStones(input []byte) []Stone {
	tokens := bytes.Fields(input)
	stones := make([]Stone, len(tokens))
	for i, token := range tokens {
		stones[i] = stoneFromStr(token)
	}
	return stones
}

func countAfter(stones []Stone, blinks int) int {
	total := 0
	for _, stone := range stones {
		total += countAfterBlinks(stone, blinks)
	}
	return total
}

type BlinkCache map[int]int

var bigBlinkCache = make(map[Stone]BlinkCache)

func countAfterBlinks(stone Stone, numBlinks int) (count int) {
	if numBlinks == 0 {
		return 1
	}
	blinkCache, ok := bigBlinkCache[stone]
	if !ok {
		blinkCache = make(BlinkCache)
		bigBlinkCache[stone] = blinkCache
	} else {
		cachedCount, isCached := blinkCache[numBlinks]
		if isCached {
			return cachedCount
		}
	}
	defer func() { blinkCache[numBlinks] = count }()

	left, right := stone.Blink()
	count += countAfterBlinks(left, numBlinks-1)
	if !right.IsNull() {
		count += countAfterBlinks(right, numBlinks-1)
	}
	return count
}
