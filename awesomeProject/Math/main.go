package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Shantik")
	{

		var x int

		fmt.Print("how old are you?1: ")
		fmt.Fscan(os.Stdin, &x)

		fmt.Println("Im  ", x)
	}
	{
		var x string
		var y int
		fmt.Printf("what is your name? How old are you?")
		fmt.Fscan(os.Stdin, &x, &y)
		fmt.Printf("My name is %v, Im %v", x, y)

	}
	{
		var x string
		var y string
		var c string
		fmt.Printf("What is your name?")
		fmt.Fscan(os.Stdin, &x)
		fmt.Printf("Name is %v\n", x)
		fmt.Printf("What is your last name?")
		fmt.Fscan(os.Stdin, &y)
		fmt.Printf("Last name is %v\n", y)
		fmt.Printf("What is your gender?")
		fmt.Fscan(os.Stdin, &c)
		fmt.Printf("Gender is %v\n ", c)
	}

	{

		c := subtraction(10, 2)
		fmt.Printf("%v", c)
	}
	{
		c := addition(1, 53)
		fmt.Printf("%v", c)
	}
	{
		c := division(25, 5)
		fmt.Printf("%v", c)
	}
	//возращение функции стринг типа данных
	{
		var Shantik string = GetSting()
		fmt.Println(Shantik)
	}
	{
		//возращение функции интов типа данных
		var (
			//Name  string = "My name is Shantal "
			myage int = GetInt()
		)

		fmt.Println(myage) //GetInt
		//fmt.Println(Name)  // My name is Shantal
	}
	//случайные цифры
	{
		// Getting result from intrandom () function
		res1 := intrandom(2, 6)
		res2 := intrandom(49, 50)
		res3 := intrandom(120, 98)

		// Dicplaying results
		fmt.Println("Result 1: ", res1)
		fmt.Println("Result 2: ", res2)
		fmt.Println("Result 3: ", res3)
	}
	//циклы
	{
		var i int
		for i = 0; i >= 0; i++ {
			fmt.Println(i)
		}

	}
	{
		var x, y int = 3, 4
		result := add(float64(x), float64(y))
		fmt.Print(result)

	}

}
