package main

import (
	"awesomeProject17/Arrays/find"
	"fmt"
)

func main() {

	arr := []int{3, 15, 0, 9, 21}

	min := arr[0]
	max := arr[0]
	sum := 0
	for i := 0; i < len(arr); i++ {
		if min < arr[i] {
			min = arr[i]

		}
		if max > arr[i] {
			max = arr[i]

		}
		sum = sum + arr[i]
	}
	avg := sum / len(arr)
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(avg)

	{
		Num := 10
		increment()
		fmt.Println(Num)

	}
	{
		Array := []int{1, 2, 3, 5, 9}
		Result := 0
		Sum1(Array, &Result)
		fmt.Println(Result)

	}

	{
		Massiv := [10]int{1, 1, 2, 1, 3, 1, 4, 1}
		fmt.Print(Massiv)
		for i := 0; i < len(arr); i++ {
		}
	}
	a := 16
	b := 44

	if a < b {
		fmt.Println("min ", a)
		fmt.Println("max ", b)

	} else {
		fmt.Println("min", b)
		fmt.Println("max ", a)
	}
	f := 35
	d := 66
	c := 2
	n := 74
	m := (f + d + c + n) / 4
	fmt.Println("m", m)
	//arr := []int{3, 15, 2, 9, 21}

	//min := arr[0]
	//for _, val := range arr {
	//	if min > val {
	//		min = val
	//	}
	//}
	//
	//fmt.Println(min)
	//max := arr[0]
	//for _, val := range arr {
	//	if max < val {
	//		max = val
	//	}

	//}
	//Увеличение данных в массиве на их количество
	{

		arr := []int{3, 5, 6, 7, 8}
		for i := 0; i < len(arr); i++ {
			fmt.Println("i", i)
			arr[i] = arr[i] + len(arr)
			fmt.Println(arr)
			fmt.Println(i)
		}
	}
	//умножений четных чисел
	{
		arr := []int{1, 2, 3, 4, 5, 6}

		for i := 0; i < len(arr); i++ {
			if i%2 != 0 {
				arr[i] = arr[i] * arr[i]

			}
		}

		fmt.Println(arr)

	}

	{
		arr := []int{1, 2, 3, 4}
		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 { // 4
				arr[i] = arr[0]
				continue
			}
			arr[i] = arr[i+1]
		}
		fmt.Println(arr)
	}
	//сложение двух массивов
	{
		a := [4]int{1, 2, 3, 4}
		b := [4]int{1, 2, 3, 4}
		c := [4]int{}
		for i := 0; i < len(a); i++ {

			c[i] = a[i] + b[i]
		}

		fmt.Println(c)
	}
	//{
	//	inc := increment
	//	print(inc)
	//
	//}
	//сложение двух массивов, четные числа сложить не четные вычесть
	{
		a := [5]int{1, 2, 3, 4, 5}
		b := [5]int{1, 2, 3, 4, 5}
		c := [5]int{}
		for i := 0; i < len(a); i++ {
			if i%2 != 0 {
				c[i] = a[i] + b[i]
			} else {
				c[i] = a[i] - b[i]
			}

		}
		fmt.Println(c)
	}

	{
		var result int
		x := []int{33, 3, 3, 3, 3, 4, 13, 5}
		result = find.Find(x, 3)
		fmt.Println(x)
		fmt.Print(result)
		{
			Massiv := [10]int{1, 1, 2, 1, 3, 1, 4, 1}
			fmt.Print(Massiv)
			for i := 0; i < len(Massiv); i++ {
			}
		}

	}
}
