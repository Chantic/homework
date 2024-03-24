package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	b := [4]int{1, 2, 3, 4}
	c := [4]int{}
	for i := 0; i < len(a); i++ {

		c[i] = a[i] + b[i]
	}

	fmt.Println(c)
}
