package main

import "fmt"

func main() {

	//задача 1
	// найти максимальный и минимальный элемент в массиве

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	min := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] < min {
			min = in[i]
		}
	}

	max := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}

	fmt.Printf("min: %d, max: %d\n", min, max)
	fmt.Print()

	ini := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var mini int
	var maxi int
	mini = ini[0]
	maxi = ini[0]
	for i := 0; i < len(ini); i++ {
		if ini[i] < min {
			min = ini[i]
		} else {
			max = ini[i]
		}

	}
	fmt.Printf("min: %d, max: %d\n", mini, maxi)

}
