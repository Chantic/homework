package main

import "fmt"

func main() {
	array := [4][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	//эта штука для записывает в двух мерный массив
	//обход массива с конца
	sum := 1
	for i := len(array) - 1; i >= 0; i-- {
		for j := len(array[i]) - 1; j >= 0; j-- {
			array[i][j] = sum
			sum = sum + 1
		}

	}

	//это штука чтобы вывести массив двух мерный
	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j], " ")
		} //
		fmt.Println("\n")
	}
}
