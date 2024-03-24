package main

import "fmt"

type Car struct {
	Brand  string
	Number int
	Size   int
	Code   string
	Iswork bool
}

func main() {
	var Nissan Car = Car{
		Brand:  "Nissan",
		Number: 292203,
		Size:   34093958,
		Code:   "fdhdjks",
	}
	var Nissan2 Car = Car{
		Brand:  "Nissan",
		Number: 292203,
		Size:   34093958,
		Code:   "fdhdjks",
	}
	var Toyota Car = Car{
		Brand:  "Toyota",
		Number: 292203,
		Size:   34093958,
		Code:   "fdhdjks",
	}
	var Toyota2 Car = Car{
		Brand:  "Toyota",
		Number: 292203,
		Size:   34093958,
		Code:   "fdhdjks",
	}

	var Cars []Car = []Car{
		Nissan, Nissan2, Toyota, Toyota2,
	}

	for i, car := range Cars {
		if car.Brand == "Toyota" {
			car.Size = 21
			Cars[i] = car
		} else {
			car.Size = 31
			Cars[i] = car
		}

	}

	for _, car := range Cars {
		fmt.Printf("c%+v\n", car)
	}
}

func Start(car *Car) {
	car.Iswork = true
}
func Stop(car *Car) {
	car.Iswork = false
}

//Nissan.Model = "Сидан"
//fmt.Printf("The Car %+v\n ", Nissan)
//
//Start(&Nissan)
//fmt.Printf("dosnt work %+v\n", Nissan)
//
//}
