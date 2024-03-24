package main

import "fmt"

func main() {
	in := "apple"

	if in == "apple" {
		fmt.Println("100")
	} else if in == "orange" {
		fmt.Println("200")
	} else if in == "banana" {
		fmt.Println("3000")
	} else {
		fmt.Println("такого продукта нет")
	}

	flag := false
	if in == "apple" {
		fmt.Println("100")
		flag = true
	}
	if in == "orange" {
		fmt.Println("200")
		flag = true
	}
	if in == "banana" {
		fmt.Println("3000")
		flag = true
	}

	if !flag {
		fmt.Println("такого продукту нет")
	}

	switch in {
	case "apple":
		fmt.Println("100")
	case "orange":
		fmt.Println("200")
	case "banana":
		fmt.Println("300")
	default:
		fmt.Println("такого продукта нет")
	}

}
