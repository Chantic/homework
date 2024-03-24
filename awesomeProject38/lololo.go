package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [64][2]string{{}, {}}

	var b rune
	b = 1072
	for i := 0; i < 32; i++ {
		arr[i][0] = string(b)
		b = b + 1
	}
	var a rune
	a = 1040
	for i := 32; i < 64; i++ {
		arr[i][0] = string(a)
		a = a + 1

	}
	// вторая часть задачи
	//запись цифр
	for i := 0; i < 32; i++ {
		var c string
		c = strconv.Itoa(i + 1)
		arr[i][1] = c

	}
	var c int
	c = 0
	for i := 32; i < 64; i++ {

		k := strconv.Itoa(c + 1)
		c = c + 1
		arr[i][1] = k

	}
	fmt.Printf("%v", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print((arr[i][0]))
		fmt.Println("\n")
		fmt.Print((arr[i][1]))

	}
}

func main2() {
	arr := [64][2]string{{}, {}}

	var a rune
	var b rune
	b = 1072
	a = 1040

	for i := 0; i < 64; i++ {
		//if i < 32{
		//	arr[i][0] = string(b)
		//	b = b + 1
		//}
		//if i >32 {
		//	arr[i][0] = string(a)
		//	a = a + 1
		//}

		if i < 32 {
			arr[i][0] = string(b)
			b = b + 1
		} else {
			arr[i][0] = string(a)
			a = a + 1
		}

	}

	// вторая часть задачи
	//запись цифр
	for i := 0; i < 32; i++ {
		var c string
		c = strconv.Itoa(i + 1)
		arr[i][1] = c

	}
	var c int
	c = 0
	for i := 32; i < 64; i++ {
		k := strconv.Itoa(c + 1)
		c = c + 1
		arr[i][1] = k

	}
	fmt.Printf("%v", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print((arr[i][0]))
		fmt.Println("\n")
		fmt.Print((arr[i][1]))

	}
}
