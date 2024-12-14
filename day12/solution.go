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
	plots := plotGrid(input)
	regions := findRegions(plots)
	price := 0
	for _, region := range regions {
		area := getArea(region)
		lines := getStraightLines(region)
		price += area * lines
	}
	return price
}

func (Solution) GetExample(part int) []byte {
	return []byte("RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE")
}

func (Solution) ExampleAnswer1() int {
	return 1930
}
func (Solution) ExampleAnswer2() int {
	return 1206
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
				newRegion.Add(v)
				seen.Add(v)
				regionSize := 0
				for newRegion.Len() != regionSize {
					regionSize = newRegion.Len()
					for p := range newRegion.Iterate() {
						for _, adj := range p.GetCardinalAdjacents() {
							if utils.WithinGrid(adj, gridSize) && grid[adj.Down][adj.Right] == plot {
								newRegion.Add(adj)
								seen.Add(adj)
							}
						}
					}
				}
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

func isAdjacent(v1 utils.VectorI, v2 utils.VectorI) bool {
	dDown := utils.Abs(v1.Down - v2.Down)
	dRight := utils.Abs(v1.Right - v2.Right)
	return (dDown == 1 && dRight == 0) || (dRight == 1 && dDown == 0)
}

func getStraightLines(region utils.Set[utils.VectorI]) int {
	externalCorners := utils.NewSet[utils.VectorI]()
	internalCorners := utils.NewSet[utils.VectorI]()
	doubleCorners := utils.NewSet[utils.VectorI]()
	for plot := range region.Iterate() {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				boxCenter := utils.VectorI{Down: plot.Down + i, Right: plot.Right + j}
				box := make([]utils.VectorI, 0, 4)
				for bi := -1; bi < 1; bi++ {
					for bj := -1; bj < 1; bj++ {
						boxElem := utils.VectorI{
							Down:  plot.Down + i + bi,
							Right: plot.Right + j + bj,
						}
						if region.Contains(boxElem) {
							box = append(box, boxElem)
						}
					}
				}
				if len(box) == 3 {
					internalCorners.Add(boxCenter)
				} else if len(box) == 2 && !isAdjacent(box[0], box[1]) {
					doubleCorners.Add(boxCenter)
				} else if len(box) == 1 {
					externalCorners.Add(boxCenter)
				}
			}
		}
	}
	totalCorners := internalCorners.Len() + externalCorners.Len() + (2 * doubleCorners.Len())
	return totalCorners
}
