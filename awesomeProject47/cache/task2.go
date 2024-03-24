package main

import (
	"fmt"
	time "time"
)

var cache1 = make(map[int]int)

type One struct {
	o map[int]int
}

func main() {
	Nat := One{
		o: map[int]int{},
	}
	lol := time.Now()
	end := someFunc(10, &Nat)
	operation_Time := lol.Second() - time.Now().Second()
	fmt.Printf("result %v, time %v", end, operation_Time)
	key, exist := Nat.o[10]
	if exist {
		fmt.Println(key)
	} else {
		end := someFunc(10, &Nat)
		Nat.o[10] = end
		fmt.Println(end)
	}
}
func someFunc(key int, cache1 *One) int {

	exist := cache1.Exist2(key)
	if !exist {
		cache1.Set2(key, key)
		return key
	}
	value, _ := cache1.Get2(key)

	return value
}
func someLogic(val int) int {
	time.Sleep(time.Second * 5)
	return val * 3
}
func (ch *One) Exist2(key int) bool {
	_, exist := ch.o[key]
	return exist
}
func (ch *One) Set2(key, value int) bool {
	ch.o[key] = value
	return true
}
func (ch *One) Delete(key int) bool {
	delete(ch.o, key)
	return true
}
func (ch *One) Get2(key int) (int, bool) {
	key, exist := ch.o[key]
	if exist {
		return key, true
	}
	return key, false
}
