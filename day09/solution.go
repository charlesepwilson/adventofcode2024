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
	ints := toIntegerList(input)
	blocks := toBlocks(ints)
	return computeChecksum2(blocks)
}

func (Solution) GetExample() []byte {
	return []byte("2333133121414131402")
}

func (Solution) ExampleAnswer1() int {
	return 1928
}
func (Solution) ExampleAnswer2() int {
	return 2858
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
	return lastBlock.id, blocks
}

func computeChecksum(blocks []Block) int {
	result := 0
	i := 0
	modifiedBlocks := make([]Block, len(blocks))
	copy(modifiedBlocks, blocks)
	totalSize := 0
	for _, block := range blocks {
		totalSize += block.size
	}
	for blockIndex := 0; i < totalSize; blockIndex++ {
		block := modifiedBlocks[blockIndex]
		for j := 0; j < block.size && i < totalSize; j++ {
			result += block.id * i
			i++
		}
		for j := 0; j < block.padding && i < totalSize; j++ {
			stolenValue, b := stealValue(modifiedBlocks)
			modifiedBlocks = b
			result += stolenValue * i
			i++
		}
	}
	return result
}

func totalLength(blocks []Block) int {
	result := 0
	for _, block := range blocks {
		result += block.size
		result += block.padding
	}
	return result
}

func defragIteration(blocks []Block, targetId int, smallestFailure int) ([]Block, int, int) {
	targetIndex := -1
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].id == targetId {
			targetIndex = i
			break
		}
	}
	if blocks[targetIndex].size >= smallestFailure {
		return blocks, smallestFailure, targetIndex
	}
	for j := 0; j < targetIndex; j++ {
		if blocks[j].padding >= blocks[targetIndex].size {
			newBlocks := make([]Block, 0, len(blocks))
			newBlocks = append(newBlocks, blocks[:j]...)
			acceptorBlock := blocks[j]
			targetBlock := blocks[targetIndex]
			preTargetBlock := blocks[targetIndex-1]
			newBlocks = append(
				newBlocks,
				Block{
					id:      acceptorBlock.id,
					size:    acceptorBlock.size,
					padding: 0,
				},
			)
			if (targetIndex - j) > 1 {
				newBlocks = append(
					newBlocks,
					Block{
						id:      targetBlock.id,
						size:    targetBlock.size,
						padding: acceptorBlock.padding - targetBlock.size,
					},
				)
				newBlocks = append(newBlocks, blocks[j+1:targetIndex-1]...)
				newBlocks = append(
					newBlocks,
					Block{
						id:      preTargetBlock.id,
						size:    preTargetBlock.size,
						padding: preTargetBlock.padding + targetBlock.padding + targetBlock.size,
					},
				)
			} else {
				newBlocks = append(
					newBlocks,
					Block{
						id:      targetBlock.id,
						size:    targetBlock.size,
						padding: targetBlock.padding + acceptorBlock.padding,
					},
				)
			}
			newBlocks = append(newBlocks, blocks[targetIndex+1:]...)
			return newBlocks, smallestFailure, j + 1
		}
	}
	if blocks[targetIndex].size < smallestFailure {
		smallestFailure = blocks[targetIndex].size
	}
	return blocks, smallestFailure, targetIndex
}

func computeChecksum2(blocks []Block) int {
	smallestFailure := 10 // maximum block size is 9 since single characters are used to define block sizes
	total := 0
	var newIndex int
	for id := len(blocks) - 1; id >= 0; id-- {
		blocks, smallestFailure, newIndex = defragIteration(blocks, id, smallestFailure)
		leadUp := totalLength(blocks[:newIndex])
		// once a block has been placed, it will never actually move its position (in the expanded representation)
		// so we can immediately compute its contribution to the checksum
		justPlacedBlock := blocks[newIndex]
		total += id * ((justPlacedBlock.size * leadUp) + (justPlacedBlock.size-1)*justPlacedBlock.size/2)
	}
	return total
}
