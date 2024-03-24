package main

import "fmt"

func main() {
	c := addition(1, 53)
	fmt.Printf("%v", c)
}
func addition(a int, b int) int {
	return a + b
}
