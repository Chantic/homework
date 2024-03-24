package main

import "fmt"

func SRectange() float64 {
	var width, height float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&width, &height)
	area := width * height
	fmt.Printf("Площадь фигуры равна:%.2f\n", area)
	return 0
}
