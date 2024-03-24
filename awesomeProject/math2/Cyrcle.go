package main

import "fmt"

func main() {
	var X int
	var Y int
	fmt.Print("Входит ли точка в круг")
	fmt.Print(" введи X ")
	fmt.Scan(&X)
	fmt.Print(" введи Y ")
	fmt.Scan(&Y)
	x1 := -5
	y1 := 5
	x2 := 5
	y2 := -5
	if X >= x1 || X <= x2 && X <= y1 || X >= y2 {
		fmt.Print("Входит  в круг")
	} else {
		fmt.Print("не входит  в круг")
	}
	{
		var X int
		var Y int
		var S int
		fmt.Print("Входит ли точка в круг\n")
		fmt.Print(" введи X ")
		fmt.Scan(&X)
		fmt.Print(" введи Y ")
		fmt.Scan(&Y)
		fmt.Print(" введи S ")
		fmt.Scan(&S)
		x1 := -(S / 2)
		y1 := S / 2
		x2 := S / 2
		y2 := -(S / 2)

		if X >= x1 || X <= x2 && X <= y1 || X >= y2 {
			fmt.Print("Входит  в круг")
		} else {
			fmt.Print("не входит  в круг")
		}
	}
}
