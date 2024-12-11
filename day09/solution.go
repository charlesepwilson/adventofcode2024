package day09

import (
	"fmt"
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
	fmt.Println(len(input))
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

func totalLength(blocks []Block) int {
	result := 0
	for _, block := range blocks {
		result += block.size
		result += block.padding
	}
	return result
}

func defragIteration(blocks []Block, targetId int, smallestFailure int) ([]Block, int) {
	targetIndex := -1
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].id == targetId {
			targetIndex = i
			break
		}
	}
	if blocks[targetIndex].size >= smallestFailure {
		fmt.Println("skipping check for", targetId, "as size is", blocks[targetIndex].size, "and already had a failure of size", smallestFailure)
		return blocks, smallestFailure
	}
	for j := 0; j < (targetIndex - 1); j++ {
		if blocks[j].padding >= blocks[targetIndex].size {
			fmt.Println("Moving id", targetId, "with size", blocks[targetIndex].size, "from index", targetIndex, "to padding of index", j, "with id", blocks[j].id, "and padding", blocks[j].padding)
			//fmt.Println(targetId, j, targetIndex)
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
			newBlocks = append(newBlocks, blocks[targetIndex+1:]...)
			if len(newBlocks) != len(blocks) {
				fmtStr := fmt.Sprintf("%d should be %d", len(newBlocks), len(blocks))
				panic(fmtStr)
			}
			totOld := totalLength(blocks)
			totNew := totalLength(newBlocks)
			if totOld != totNew {
				fmtStr := fmt.Sprintf("%d should be %d", totNew, totOld)
				panic(fmtStr)
			}
			return newBlocks, smallestFailure
		}
	}
	fmt.Println("couldn't find place for block", targetId, "with size", blocks[targetIndex].size)
	if blocks[targetIndex].size < smallestFailure {
		smallestFailure = blocks[targetIndex].size
	}
	return blocks, smallestFailure
}

func toString(blocks []Block) string {
	result := make([]byte, 0)
	for _, block := range blocks {
		for j := 0; j < block.size; j++ {
			b := []byte(strconv.Itoa(block.id))
			if len(b) == 1 {
				result = append(result, b[0])
			} else {
				result = append(result, '(')
				result = append(result, b...)
				result = append(result, ')')
			}
		}
		for j := 0; j < block.padding; j++ {
			result = append(result, '.')
		}
	}
	return string(result)
}

func printBlocks(blocks []Block) {
	//fmt.Println(toString(blocks))
}

func computeChecksum2(blocks []Block) int {
	fmt.Println(len(blocks))
	smallestFailure := 10
	for b := len(blocks) - 1; b >= 0 && smallestFailure > 1; b-- {
		printBlocks(blocks)
		blocks, smallestFailure = defragIteration(blocks, b, smallestFailure)
	}
	printBlocks(blocks)
	result := 0
	i := 0
	for _, block := range blocks {
		for j := 0; j < block.size; j++ {
			result += block.id * i
			i++
		}
		for j := 0; j < block.padding; j++ {
			i++
		}
	}
	// 6357408289030 too high
	// 6353648838485 too high
	// 6353625164652 too low
	return result
}
