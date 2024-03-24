package main

import "fmt"

func main() {
	year := 400
	if isLeapYear(year) {
		fmt.Printf("%d является высокосным годом", year)
	} else {
		fmt.Printf("%d не является высокосным годом", year)
	}
	{

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

}
