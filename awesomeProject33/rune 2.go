package main

import "fmt"

func main() {
	array := [4][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	sum := 1
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			array[i][j] = sum
			sum = sum + 1
		}

	}

	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j], " ")
		}
		fmt.Println("\n")
	}

}
