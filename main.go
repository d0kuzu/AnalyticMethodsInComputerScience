package main

import (
	"AnalyticMethodsInComputerScience/practice2/Josephus"
	"fmt"
)

func main() {
	fmt.Println("Josephus")
	for i := 1; i <= 7; i++ {
		recursiveResult := Josephus.CalculateRecursive(i, 3)
		fmt.Printf("%d (recursive {%d}) \n", recursiveResult, i)
	}
}
