package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{}
	for i := 0; i < len(a); i++ {
		if i%2 != 0 {
			c[i] = a[i] + b[i]
		} else {
			c[i] = a[i] - b[i]
		}

	}
	fmt.Println(c)
}
