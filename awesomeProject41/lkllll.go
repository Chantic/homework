package main

import "fmt"

func double(x *int) {
	//x copy 0x123
	*x += *x
	x = nil // the line is just for explanation purpose
}

func main() {
	var a = 3
	double(&a)               // a =6
	fmt.Println(a)           // 6
	p := &a                  // p = 0x32
	double(p)                // a =12
	fmt.Println(a, p == nil) // 12 false
}
