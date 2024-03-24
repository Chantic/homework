package main

import (
	"fmt"
	"math"
)

func main() {

	var choice string

	fmt.Println("Выбери фигуру:\n a)Прямоугольник\n  b)Треугольник\n  c)Круг")
	fmt.Scan(&choice)

	switch choice {
	case "Прямоугольник":
		fmt.Println(SRectange())
	case "Треугольник":

		fmt.Println(STriange())
	case "Круг":
		fmt.Println(SCircle())
	default:
		fmt.Println("Некорректный ввод.")
	}
}
func SRectange() float64 {
	var width, height float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&width, &height)
	area := width * height
	fmt.Printf("Площадь фигуры равна:%.2f\n", area)
	return 0
}
func STriange() float64 {
	var base, height float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&base, &height)
	area := (base * height) / 2
	fmt.Printf("Площадь фигуры равна:%.2f\n", area)
	return 0
}
func SCircle() float64 {
	var radius float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&radius)
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("SCircle:%.2f\n", area)
	return 0
}
