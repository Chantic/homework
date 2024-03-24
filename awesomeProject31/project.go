package main

import "fmt"

func main() {
	array := [4][4]rune{{'o', 'o', 'o', 'o'}, {'o', 'o', 'o', 'o'}, {'o', 'o', 'o', 'o'}, {'o', 'o', 'o', 'o'}}
	for i := 0; i < len(array); i++ {
		{
			for j := 0; j < len(array); j++ {
				if i%2 != 0 {
					array[i][j] = 'x'
				}
			}
		}

	}
	fmt.Print(array)
}
