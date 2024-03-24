//

package main

import (
	"fmt"
	"math/rand"
)

// Function
func intrandom(value_1, value_2 int) int {
	return value_1 + value_2 + rand.Int()
}

// Main function
func main() {
	// Getting result from intrandom () function
	res1 := intrandom(2, 6)
	res2 := intrandom(49, 50)
	res3 := intrandom(120, 98)

	// Dicplaying results
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
}
 