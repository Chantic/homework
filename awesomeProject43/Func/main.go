package main

import "fmt"

func main() {
	c := Suum(4, 4)
	fmt.Println(c)
	sayEnd := func() {
		fmt.Println("end")
	}
	Sayer(SayHi)
	Sayer(SayHiShanti)
	Sayer(sayEnd)
	a := func() int {
		return 1
	}
	i := 3
	someFunc := func(c int) {

		c++
		fmt.Println(c)
	}
	someFunc(i)
	someFunc(i)
	someFunc(i)
	fmt.Println(Summator(a))
	fmt.Println(GetOne())
	fmt.Println(makeAdder(2))
	fmt.Println(makeAdder(4))
	fmt.Println(makeAdder(3))
	fmt.Println(Get10(1))
	OnetoN(5)
	fmt.Println(FuckIlya(7))
	fmt.Println(BanIlya(10))
}

func Suum(a, b int) int {
	return a + b
}
func SayHi() {
	fmt.Println("Hi")
}
func SayHiShanti() {
	fmt.Println("Hi Shanti")
}
func Sayer(say func()) {
	say()
}
func GetOne() int {
	return 1

}
func Summator(C func() int) int {
	C()
	result := C() + 1
	return result
}
func makeAdder(x int) int {
	return func(y int) int {
		return x + y
	}(x)

}
func Get10(i int) int {
	if i >= 10 {
		return i
	}
	i++
	return Get10(i)
}
func OnetoN(n int) {
	fmt.Println("some:", n)
	if n <= 1 {
		return
	}
	n--
	OnetoN(n)

}
func FuckIlya(n int) int {

	if n <= 1 {
		return n
	}

	return FuckIlya(n-1) * n

}
func BanIlya(n int) int {
	if n <= 1 {
		return n
	}
	return BanIlya(n-1) + n
}

