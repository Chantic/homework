package Flags

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.Float64("Н1", 0, "Первое число")
	b := flag.Float64("Н2", 0, "Второе  число")
	Oper := flag.String("Оп", "*", "Опреации (+,-,*,/) ")
	flag.Parse()
	var result float64
	//fmt.Print("Введите первый аргумент  ")
	//fmt.Scan(&a)
	//fmt.Print("Введите второй  аргумент  ")
	//fmt.Scan(&b)
	//fmt.Print("Выберите операцию (+,-,*,/)")
	//fmt.Scan(&result)
	switch *Oper {
	case "+":
		result = *a + *b
	case "-":
		result = *a - *b
	case "*":
		result = *a * *b
	case "/":
		if *b == 0 {
			fmt.Println("Невозможная операция:Ибо все решили что ноль это ничего и делить его нельзя ")
			return

		}
		result = *a / *b
	default:
		fmt.Println("Невозможная операция:Ошибка ")
		return
	}
	fmt.Printf("%f %s %f = %f\n", *a, *Oper, *b, result)
}

//	type Calculator  struct{
//		m map{int }string
//
// //	}
//
//	func parceArgs(c []string) (float64, float64, error) {
//		a, err := strconv.ParseFloat(c[0], 64)
//		if err != nil {
//			return 0.0, 0.0, err
//		}
//		b, err := strconv.ParseFloat(c[2], 64)
//		{
//			return 0.0, 0.0, nil
//		}
//		return a, b, nil
//
// }
//
//	func processStrack(e []string) (float64, error) {
//		result := 0.0
//		for _, v := range e {
//			c := strings.Split(v, "")
//			if len(c)-1 < 2 {
//				return 0.0, errors.New(" Ошибка выдачи аргументов ")
//			}
//			a, b, err := parceArgs(c)
//			if err != nil {
//				return 0.0, err
//			}
//			switch c[1] {
//			case "*":
//				if b == 0.0 {
//					result = a * b
//				}
//			case "/":
//				if b == 0.0 {
//					return 0.0, errors.New("Ошиька вы пытаетесь функционировать с нулем ")
//
//				}
//				result = a / b
//			case "+":
//				result = a + b
//			case "-":
//				result = a - b
//			}
//		}
//		return result, nil
//	}
//
//	func main() {
//		expressions := make([]string, 1)
//		for {
//			scanner := bufio.NewScanner(os.Stdin)
//			fmt.Print("gocalc>")
//			for scanner.Scan() {
//				expressions = append(expressions, scanner.Text())
//				result, err := processStrack(expressions)
//				if err != nil {
//					fmt.Println(err)
//				} else {
//					fmt.Println(result)
//				}
//				fmt.Print("gocalc>")
//			}
//		}
//	}
