package main

import (
	"fmt"
	"os"
)

func main() {

	var x int

	fmt.Print("how old are you?1: ")
	fmt.Fscan(os.Stdin, &x)

	fmt.Println("Im  ", x)
}
