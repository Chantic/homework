package main

func main() {
	inc := increment
	print(inc)

}
func increment() int {
	count := 0
	//func() int {
	count += 1
	return count
}
