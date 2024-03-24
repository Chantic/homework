package main

func main() {
	arr := []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i+1]

	}
}
