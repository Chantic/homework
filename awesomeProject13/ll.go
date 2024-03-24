package main

import "fmt"

func main() {
	Massiv := [6]int{-13, 21, 6, -6, 0, 34}
	for i := 0; i < len(Massiv); i++ {
		if Massiv[i] >= 0 {
			fmt.Print(" i = ", Massiv[i])
			fmt.Println(" [", i, "] =", Massiv[i])
		}

	}
}
