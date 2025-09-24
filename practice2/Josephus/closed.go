package Josephus

import (
	"math"
)

func CalculateClosed(n int) int {
	m := int(math.Floor(math.Log2(float64(n))))
	l := n - int(math.Pow(2, float64(m)))
	return 2*l + 1

}
