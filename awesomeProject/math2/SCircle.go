package main

import (
	"fmt"
	"math"
)

func SCircle() float64 {
	var radius float64
	fmt.Println("введите данные фигуры :")
	fmt.Scan(&radius)
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("SCircle:%.2f\n", area)
	return 0
}
