package main

import (
	"fmt"
	"time"
)

var cashe = make(map[int64]int64)

type Cache struct {
	m map[int]int
}

// идея
// берем инфу
// делаем работу
// и сохраняем этот результат
func main() {
	//текущее время
	cache := Cache{
		m: map[int]int{},
	}
	now := time.Now()
	// проходим операцию
	result := LongOperationWithCache(10, &cache)
	// вычесление за сколько мы выполняем операцию
	t1 := now.Second() - time.Now().Second()
	// принтим результат
	fmt.Printf("result %v, time %v ", result, t1)
	//текущее время
	now = time.Now()
	//проходим по операции с тем же значением
	result = LongOperationWithCache(10, &cache)
	// вычесляем за сколько времени проходит
	t2 := time.Now().Second() - now.Second()
	//принтим результат
	fmt.Printf("result %v, time %v ", result, t2)
	//cache := &Cache{}
	//cache.m["key1"] = "value1"
	//cache.m["key2"] = "value2"
	//exist := cashe.Exist("key 1")
}

func LongOperation(s int64) int64 {
	//пауза на 2 секунды
	time.Sleep(time.Second * 2)
	return s / 2
}

func LongOperationWithCache(s int64, cache *Cache) int64 {
	//проверка в кэше

	var exist bool
	exist = cache.Exist(int(s))
	if !exist {
		// если нет в кэше, то выполняем что то
		result := LongOperation(s)
		cache.Set(int(s), int(result))
		return result
	}
	value, key := cache.Get(int(s))

	fmt.Println(key)

	return int64(value)

}

//получение
//удаление
//вставка
//проверка

//func NewCache() *Cache {

//}
//func Construct() *Cache {

//}

// конструктор для класс

/*
	все функции должны быть методами типа *Cache

func (c *Cache) Set(key, value string)bool // должна возвращать ошибку данные по ключу уже есть
func (c *Cache) Exist(key string) bool // функция говорит, существует ли данные по ключу

func (c *Cache) Get(key string) string, error  // либо Get(key string) string
func (c *Cache) Del(key string) bool // возбраняет ошибку если нечего удалять
*/
func (c *Cache) Exist(key int) bool {
	_, ok := c.m[key]
	return ok
}
func (c *Cache) Set(key, value int) bool {
	c.m[key] = value
	return true
}
func (c *Cache) Del(key int) bool {
	delete(c.m, key)
	return true
}
func (c *Cache) Get(key int) (int, bool) {
	value, ok := c.m[key]
	if ok {
		return value, true
	}
	return value, false
}

func (v *Cache) Delete(key int) error {

	delete(v.m, key)
	return nil

}
