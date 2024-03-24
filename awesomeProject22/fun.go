package main

import "fmt"

func isLeapYear(year int) bool {
	if year%4 == 0  {
		{
			return true
		}
	}

	return false
}
func main() {
	year := 400
	if isLeapYear(year) {
		fmt.Printf("%d является высокосным годом", year)
	} else {
		fmt.Printf("%d не является высокосным годом", year)
	}
}
