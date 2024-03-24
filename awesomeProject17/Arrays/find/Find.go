package find

//
//func main() {
//	var result int
//	x := []int{33, 3, 3, 3, 3, 4, 13, 5}
//	result = Find(x, 3)
//
//	fmt.Print(result)
//}

func Find(arr []int, desired int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == desired {
			return i
		}

	}
	return -1
}
