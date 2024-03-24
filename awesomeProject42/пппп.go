package main

import "fmt"

func main() {
	array := [5][5]string{{"0", "0", "0", "0", "0"}, {"0", "0", "0", "0", "0"}, {"0", "0", "0", "0", "0"}, {"0", "0", "0", "0", "0"}, {"0", "0", "0", "0", "0"}}

	for i := 0; i < len(array); i++ {

		{
			for j := 0; j < len(array); j++ {
				if array[i] == array[j] {
					array[i][j] = "x"
				}
				if array[i] == array[j] {
					array[i][j] = "/"
				}
				if array[i] == array[j] {
					array[i][j] = "\\"
				}
			}

		}

	}
	for _, value := range array {
		fmt.Println(value)
	}

}
