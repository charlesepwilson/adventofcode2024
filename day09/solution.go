package day09

import (
	"strconv"
)

type Solution struct{}

func (Solution) Day() int { return 9 }

func (Solution) Part1(input []byte) int {
	ints := toIntegerList(input)
	blocks := toBlocks(ints)
	return computeChecksum(blocks)
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample() []byte {
	return []byte("2333133121414131402")
}

func (Solution) ExampleAnswer1() int {
	return 1928
}
func (Solution) ExampleAnswer2() int {
	return 0
}

func toIntegerList(input []byte) []int {
	result := make([]int, len(input))
	for i, c := range input {
		integer, _ := strconv.Atoi(string([]byte{c}))
		result[i] = integer
	}
	return result
}

type Block struct {
	id, size, padding int
}

func toBlocks(ints []int) []Block {
	ints = append(ints, 0)
	result := make([]Block, len(ints)/2)
	for i := 0; i < len(ints); i += 2 {
		fileLength := ints[i]
		paddingLength := ints[i+1]
		id := i / 2
		result[id] = Block{
			id:      id,
			size:    fileLength,
			padding: paddingLength,
		}
	}
	return result
}

func stealValue(blocks []Block) (int, []Block) {
	lastBlock := blocks[len(blocks)-1]
	for lastBlock.size == 0 {
		blocks = blocks[:len(blocks)-1]
		lastBlock = blocks[len(blocks)-1]
	}
	blocks[len(blocks)-1] = Block{
		id:      lastBlock.id,
		size:    lastBlock.size - 1,
		padding: lastBlock.padding,
	}
	//fmt.Println(lastBlock.id, blocks)
	return lastBlock.id, blocks
}

func computeChecksum(blocks []Block) int {
	result := 0
	i := 0
	modifiedBlocks := make([]Block, len(blocks))
	copy(modifiedBlocks, blocks)
	//fmt.Println(modifiedBlocks)
	totalSize := 0
	for _, block := range blocks {
		totalSize += block.size
	}
	for blockIndex := 0; i < totalSize; blockIndex++ {
		block := modifiedBlocks[blockIndex]
		for j := 0; j < block.size && i < totalSize; j++ {
			result += block.id * i
			//fmt.Println(i, block.id, result)
			i++
		}
		for j := 0; j < block.padding && i < totalSize; j++ {
			stolenValue, b := stealValue(modifiedBlocks)
			modifiedBlocks = b
			result += stolenValue * i
			//fmt.Println(i, stolenValue, result)
			i++
		}
	}
	return result
}
