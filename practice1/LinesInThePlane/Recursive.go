package LinesInThePlane

func CalculateRecursive(n int) int {
	lastN := 1 //L0 = 1

	for i := 1; i <= n; i++ {
		lastN = lastN + i
	}

	return lastN
}
