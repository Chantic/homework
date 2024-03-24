package main

import "fmt"

func main() {

	arr := []int{3, 5, 6, 7, 8}
	for i := 0; i < len(arr); i++ {
		fmt.Println("i", i)
		arr[i] = arr[i] + len(arr)
		fmt.Println(arr)
		fmt.Println(i)
	}
}
