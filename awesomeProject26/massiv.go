package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 { // 4
			arr[i] = arr[0]
			continue
		}
		arr[i] = arr[i+1]
	}
	fmt.Println(arr)
}
