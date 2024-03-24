package main

import (
	"fmt"
	"os"
)

func main() {
	var x string
	var y string
	var c string
	fmt.Printf("What is your name?")
	fmt.Fscan(os.Stdin, &x)
	fmt.Printf("Name is %v\n", x)
	fmt.Printf("What is your last name?")
	fmt.Fscan(os.Stdin, &y)
	fmt.Printf("Last name is %v\n", y)
	fmt.Printf("What is your gender?")
	fmt.Fscan(os.Stdin, &c)
	fmt.Printf("Gender is %v\n ", c)
}
