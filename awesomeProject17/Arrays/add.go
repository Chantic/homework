package main

func Sum1(Array []int, Result *int) {
	s := 0
	for i := 0; i < len(Array); i++ {
		{
			s = s + Array[i]
		}
	}
	*Result = s
}
