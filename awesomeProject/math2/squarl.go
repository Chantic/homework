package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("реши квад уравненене")
	fmt.Println("введи а:")
	fmt.Scan(&a)
	fmt.Println("реши квад уравненене")
	fmt.Println("введи b:")
	fmt.Scan(&b)
	fmt.Println("реши квад уравненене")
	fmt.Println("введи c:")
	fmt.Scan(&c)
	D := (b * b) - 4*(a*c)
	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println(" Уравнение имеет два корня \nD=" + fmt.Sprint(D))
		fmt.Println("X1:" + fmt.Sprint(x1) + "\n X2" + fmt.Sprint(x2))
	}
	if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Уравнение имеет один корень\n D")
		fmt.Println(" X;+ " + fmt.Sprint(x))
	}
	if D < 0 {
		fmt.Println(" уравнение не имеет корней\nD<0\nD" + fmt.Sprint(D))
	}
}
