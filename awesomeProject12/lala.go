package main

import "fmt"

func main() {
	var x, y int = 3, 4
	result := add(float64(x), float64(y))
	fmt.Print(result)

}
func add(a, b float64) float64 {
	return a + b

}
