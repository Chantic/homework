package main

import (
	"flag"
	"fmt"
)

func main() {
	var a = flag.String("name", "Вася", "First Name")

	var b = flag.String("surname", "Пупкин", "Last Name")
	flag.Parse()
	fmt.Println("Hi,", *a, *b)
}
