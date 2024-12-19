package utils

import "strings"

type Grid [][]byte

const OUTSIDE = ';'

func (g *Grid) Get(v VectorI) byte {
	if !WithinGridSize(v, g.Size()) {
		return OUTSIDE
	}
	return (*g)[v.Down][v.Right]
}

func (g *Grid) Set(v VectorI, c byte) {
	if !WithinGridSize(v, g.Size()) {
		return
	}
	(*g)[v.Down][v.Right] = c
}

func (g *Grid) Size() VectorI {
	return VectorI{
		Down:  len(*g),
		Right: len((*g)[0]),
	}
}

func (g *Grid) String() string {
	builder := strings.Builder{}
	for _, row := range *g {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}
	return builder.String()
}

func WithinGridSize(location VectorI, gridSize VectorI) bool {
	return location.Right >= 0 && location.Down >= 0 && location.Right < gridSize.Right && location.Down < gridSize.Down
}

func (g *Grid) WithinGrid(location VectorI) bool {
	gridSize := g.Size()
	return WithinGridSize(location, gridSize)
}

var Directions = [4]VectorI{
	{Down: -1},
	{Right: 1},
	{Down: 1},
	{Right: -1},
}

func MakeGrid(size VectorI) Grid {
	grid := make(Grid, size.Down)
	for i := range grid {
		grid[i] = make([]byte, size.Right)
	}
	return grid
}
