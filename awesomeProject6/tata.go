package main

import (
	"fmt"
	"os"
)

func main() {
	var x string
	var y int
	fmt.Printf("what is your name? How old are you?")
	fmt.Fscan(os.Stdin, &x, &y)
	fmt.Printf("My name is %v, Im %v", x, y)

}
