package Josephus

func CalculateRecursive(n, k int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	index := 0
	for len(arr) > 1 {
		index = (index + k - 1) % len(arr)
		arr = append(arr[:index], arr[index+1:]...)
	}

	return arr[0]
}
