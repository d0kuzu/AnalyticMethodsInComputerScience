package main

import (
	"AnalyticMethodsInComputerScience/practice1/Hanoi"
	"AnalyticMethodsInComputerScience/practice1/LinesInThePlane"
	"fmt"
)

func main() {
	fmt.Println("Hanoi")
	for i := 1; i <= 5; i++ {
		recursiveResult := Hanoi.CalculateRecursive(i)
		closedResult := Hanoi.CalculateClosed(i)
		fmt.Printf("%d (recursive {%d}) \n", recursiveResult, i)
		fmt.Printf("%d (closed {%d}) \n", closedResult, i)
	}
	fmt.Println("\nLines in the plane\n")

	for i := 1; i <= 5; i++ {
		recursiveResult := LinesInThePlane.CalculateRecursive(i)
		closedResult := LinesInThePlane.CalculateClosed(i)
		fmt.Printf("%d (recursive {%d}) \n", recursiveResult, i)
		fmt.Printf("%d (closed {%d}) \n", closedResult, i)
	}
}
