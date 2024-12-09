package utils

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
