package main

import (
	"fmt"
	"time"
)

// one ->  Cache : потму что не говорящие имена переменных
type Cache struct {
	lmap map[int]int
}

func main() {
	Nat := Cache{
		lmap: map[int]int{},
	}

	start := time.Now()

	result := newOne(10, &Nat)

	resultTime := time.Now().Second() - start.Second()

	fmt.Printf("resultTime%v, result %v", resultTime, result)

	start2 := time.Now()

	result2 := newOne(10, &Nat)

	resultTime = time.Now().Second() - start2.Second()

	fmt.Printf("\n resultTime %v,  result %v", resultTime, result2)

}

func someFunc(key int, cache *Cache) int {
	exist := cache.Exist(key)
	if !exist {
		cache.Set(key, key)
		return key
	}

	value, _ := cache.Get(key)

	return value
}

func practiceFunc(val int) int {
	som := 12
	time.Sleep(time.Second * 5)
	return val * som
}
func newOne(s int, cache *Cache) int {

	ok := cache.Exist(s)
	if !ok {
		sommi := practiceFunc(s)
		cache.Set(s, sommi)
		return sommi
	}
	//return s
	val, _ := cache.Get(s)
	return val

}

func someLogic(val int) int {
	time.Sleep(time.Second * 5)
	return val * 3
}
func (c *Cache) Exist(key int) bool {
	_, exist := c.lmap[key]
	return exist
}
func (c *Cache) Set(key, value int) bool {
	c.lmap[key] = value
	return true
}
func (c *Cache) Delete(key int) bool {
	delete(c.lmap, key)
	return true
}
func (c *Cache) Get(key int) (int, bool) {
	key, exist := c.lmap[key]
	return key, exist
}
