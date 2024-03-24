package main

import "fmt"

func main() {
	//arr := []int{3, 15, 2, 9, 21}

	//min := arr[0]
	//for _, val := range arr {
	//	if min > val {
	//		min = val
	//	}
	//}
	//
	//fmt.Println(min)
	//max := arr[0]
	//for _, val := range arr {
	//	if max < val {
	//		max = val
	//	}
	//}

	a := 16
	b := 44

	if a < b {
		fmt.Println("min ", a)
		fmt.Println("max ", b)

	} else {
		fmt.Println("min", b)
		fmt.Println("max ", a)
	}
	f := 35
	d := 66
	c := 2
	n := 74
	m := (f + d + c + n) / 4
	fmt.Println("m", m)
}
