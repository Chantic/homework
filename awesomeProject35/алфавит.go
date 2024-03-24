package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := [26][2]string{{}, {}}
	var a rune
	a = 97
	for i := 0; i < len(arr); i++ {
		arr[i][0] = string(a)
		a = a + 1
	}

	for i := 0; i < len(arr); i++ {

		var c string
		c = strconv.Itoa(i)
		arr[i][1] = c
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(string(arr[i][1]))
		fmt.Println()
		fmt.Print(string(arr[i][0]))

	}

	//var a rune
	//a = 666
	//fmt.Println(string(a))
}
