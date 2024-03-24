package main

import "fmt"

func main() {
	var drink string

	fmt.Println("Выберите название напитка:\nLatte\n Cappuchino\n Matcha\n Hot chocolate\n Smoothy\n Tea")
	fmt.Scan(&drink)
	/* если номер больше семи — берём целочисленный остаток от деления номера на 7 */

	/* выбираем день из вариантов ниже */
	switch drink {
	/* если день равен 0 */
	case "Latte":
		fmt.Print("12 euro")
		/* выходим из дальнейших проверок */
		break
	case "Cappuchino":
		fmt.Print("12 euro")
		break
	case "Matcha ":
		fmt.Print("12 euro ")
		break
	case "Hot chocolate":
		fmt.Print("10 euro")
		break
	case "Smoothy":
		fmt.Print("`10 euro`")
		break
	case "Tea":
		fmt.Print("5 euro")
		break
		/* если ни один вариант не подошёл */
	default:
		fmt.Print("нет такого напитка")
	}
}
