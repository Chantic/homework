package main

import "fmt"

func STriange() float64 {
	var base, height float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&base, &height)
	area := (base * height) / 2
	fmt.Printf("Площадь фигуры равна:%.2f\n", area)
	return 0
}
