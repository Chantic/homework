package main

import "fmt"

func main() {
	var result int
	x := []int{33, 3, 3, 3, 3, 4, 13, 5}
	result = Find(x, 3)

	fmt.Print(result)
	{
		Massiv := [10]int{1, 1, 2, 1, 3, 1, 4, 1}
		fmt.Print(Massiv)
		for i := 0; i < len(Massiv); i++ {
		}
	}

}
func Find(arr []int, desired int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == desired {
			return i
		}

	}
	return -1
}
