package main

import "fmt"

type People struct {
	Age         int
	Character   string
	Nationality string
	Name        string
	LastName    string
	High        int
}

func (p *People) Birthday() {
	p.Age = p.Age + 1
}
func (p *People) Lateryear() {
	p.Age = p.Age + 1
}
func main() {
	var Ilya People = People{
		Age:         24,
		Character:   "Husky",
		Nationality: "Ukrain",
		Name:        "Ilya ",
		LastName:    "Zakharkiv",
		High:        185,
	}
	var Milana People = People{
		Age:         26,
		Character:   "Cute",
		Nationality: "Tatar",
		Name:        "Milana ",
		LastName:    "Azam-Zangegeh",
		High:        169,
	}
	var Aziz People = People{
		Age:         26,
		Character:   "annoyed",
		Nationality: "Iran",
		Name:        "Aziz ",
		LastName:    "Azam-Zangegeh",
		High:        184,
	}
	var _ People = People{
		Age:         30,
		Character:   "crippe",
		Nationality: "russian",
		Name:        "Ira ",
		LastName:    "noname",
		High:        160,
	}
	Ilya.Nationality = "russian"
	//fmt.Printf("%+v\n %+v\n %+v\n", Ilya, Milana, Aziz)
	room := []People{Milana, Aziz, Ilya}
	age := 0
	a := 0
	for i := 0; i < len(room); i++ {
		room[i].Lateryear()
	}

	for i := 0; i < len(room); i++ {
		age += room[i].Age

		if room[i].Age >= 25 {

			a++

		}

	}
	fmt.Println(a)
	Milana.Birthday()
	fmt.Println(Milana)
	fmt.Println(room)
	fmt.Println(age)
}
