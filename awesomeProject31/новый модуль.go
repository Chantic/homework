package main

import "fmt"

func main() {

	var arr [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := len(arr) - 1; i >= 0; i-- {
		for j := len(arr) - 1; j >= 0; j-- {
			fmt.Print(arr[i][j], "")
		}
	}

}

//for i := 0; i < 3; i++ {
//	for j := 0; j < 3; j++ {
//		fmt.Print(arr[i][j])
//	}
//	fmt.Println()
//}

//sum := 0
//for i := 0; i < 3; i++ {
//	for j := 0; j < 3; j++ {
//		sum = sum + arr[i][j] // 00  sum  0  // sum == 0 + 1 = 1  //  0 1  sum = 1 + 2
//	}
//}

//multiply := 1

//for i := 0; i < 3; i++ {
//for j := 0; j < 3; j++ {

//multiply = multiply * arr[i][j]
//}

//fmt.Print(multiply)
