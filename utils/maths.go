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
