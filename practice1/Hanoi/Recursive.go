package Hanoi

func CalculateRecursive(n int) int {
	lastN := 1 //N0 = 1

	for i := 1; i < n; i++ {
		lastN = 2*lastN + 1
	}

	return lastN
}
