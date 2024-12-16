package utils

import (
	"math"
)

func Abs[T Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Count[T comparable](slice []T, item T) int {
	count := 0
	for _, s := range slice {
		if s == item {
			count++
		}
	}
	return count
}

func Pow(base int, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

type VectorI struct {
	Down, Right int
}

func (v VectorI) Add(other VectorI) VectorI {
	return VectorI{v.Down + other.Down, v.Right + other.Right}
}

func (v VectorI) Sub(other VectorI) VectorI {
	return VectorI{v.Down - other.Down, v.Right - other.Right}
}

func (v VectorI) Mul(val int) VectorI {
	return VectorI{v.Down * val, v.Right * val}
}

func (v VectorI) TurnRight() VectorI {
	return VectorI{v.Right, -v.Down}
}

func Gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

//func Lcm(a, b int) int {
//	return a * b / Gcd(a, b)
//}

func (v VectorI) Simplify() VectorI {
	gcd := Gcd(Abs(v.Right), Abs(v.Down))
	return VectorI{Down: v.Down / gcd, Right: v.Right / gcd}
}

func (v VectorI) GetDiagAdjacents() []VectorI {
	adjacents := make([]VectorI, 0, 8)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			adjacents = append(adjacents, v.Add(VectorI{i, j}))
		}
	}
	return adjacents
}

func (v VectorI) GetCardinalAdjacents() []VectorI {
	adjacents := make([]VectorI, 0, 4)
	for _, delta := range [4]VectorI{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	} {
		adjacents = append(adjacents, v.Add(delta))
	}
	return adjacents
}

func AlmostEqual(a, b float64) bool {
	epsilon := 0.0001
	delta := math.Abs(a - b)
	return delta < epsilon
}

func EuclideanMod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
