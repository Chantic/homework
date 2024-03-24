package main

import "fmt"

func main() {
	//blackRuns := []rune{'a', 'p','g','A'}
	s := []rune("apple")
	for i := 0; i < len(s); i++ {
		if s[i] == rune('a') {
			s[i] = '#'
		}

	}
	fmt.Println(string(s))
}
