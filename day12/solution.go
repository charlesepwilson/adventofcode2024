package day12

import (
	"advent_of_code_2024/utils"
	"bytes"
)

type Solution struct{}

func (Solution) Day() int { return 12 }

func (Solution) Part1(input []byte) int {
	plots := plotGrid(input)
	regions := findRegions(plots)
	price := 0
	for _, region := range regions {
		price += getArea(region) * getPerimeter(region)
	}
	return price
}

func (Solution) Part2(input []byte) int {
	return len(input)
}

func (Solution) GetExample() []byte {
	return []byte("RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE")
}

func (Solution) ExampleAnswer1() int {
	return 1930
}
func (Solution) ExampleAnswer2() int {
	return 0
}

func plotGrid(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}

func findRegions(grid [][]byte) []utils.Set[utils.VectorI] {
	var regions []utils.Set[utils.VectorI]
	gridSize := utils.VectorI{Down: len(grid), Right: len(grid[0])}
	seen := utils.NewSet[utils.VectorI]()
	for i, row := range grid {
		for j, plot := range row {
			v := utils.VectorI{Down: i, Right: j}
			if !seen.Contains(v) {
				newRegion := utils.NewSet[utils.VectorI]()
				utils.ToggleablePrint("Building region of", string([]byte{plot}), "starting at", v)
				newRegion.Add(v)
				seen.Add(v)
				regionSize := 0
				for newRegion.Len() != regionSize {
					regionSize = newRegion.Len()
					for p := range newRegion.Iterate() {
						for _, adj := range p.GetCardinalAdjacents() {
							if utils.WithinGrid(adj, gridSize) && grid[adj.Down][adj.Right] == plot {
								utils.ToggleablePrint("Adding to region", adj)
								newRegion.Add(adj)
								seen.Add(adj)
							}
						}
					}
					utils.ToggleablePrint("Region Size updated from", regionSize, " to", newRegion.Len())
				}
				utils.ToggleablePrint("region", string([]byte{plot}), newRegion)
				regions = append(regions, newRegion)
			}
		}
	}

	return regions
}

func getArea(region utils.Set[utils.VectorI]) int {
	return region.Len()
}

func getPerimeter(region utils.Set[utils.VectorI]) int {
	perimeter := region.Len() * 4
	for plot := range region.Iterate() {
		for _, adjacent := range plot.GetCardinalAdjacents() {
			if region.Contains(adjacent) {
				perimeter--
			}
		}
	}
	return perimeter
}
