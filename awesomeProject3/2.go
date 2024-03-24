package main

import (
	"fmt"
)

func main() {

	var (
		//Name  string = "My name is Shantal "
		myage int = GetInt()
	)

	fmt.Println(myage) //GetInt
	//fmt.Println(Name)  // My name is Shantal
}
func GetInt() int {
	return 21
}
