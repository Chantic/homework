package main

import "fmt"

func main() {
	c := division(25, 5)
	fmt.Printf("%v", c)
}
func division(a int, b int) int {
	return a / b
}
