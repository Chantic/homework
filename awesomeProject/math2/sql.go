package main

import "fmt"

func main() {
	var a float64

	println("введи a")
	fmt.Scan(&a)
	//a*(c*2)
	if a > 0 {
		S := a * a
		fmt.Print(S)
	} else {
		fmt.Print("квадрат не существует ")
	}

}
