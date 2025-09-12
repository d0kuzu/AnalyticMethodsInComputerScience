package Hanoi

import "math"

func CalculateClosed(n int) int {
	return int(math.Pow(2, float64(n)) - 1)
}
