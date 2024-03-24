package main

import (
	"fmt"
	"strings"
)

type Translate struct {
	en  string
	kaz string
}

func main() {

	//1 завернуть эту задачу в отдельную функцию
	//2 ру англ кз
	// подсказки оператор switch
	// в качестве value можно хранить структуру
	text := "Забавные птицы – сойки-пересмешницы," +
		" зато Капитолию они точно бельмо на глазу." +
		" Когда восстали дистрикты, " +
		"для борьбы с ними в Капитолии вывели генетически измененных животных." +
		" Их называют перерождениями или просто переродками." +
		" Одним из видов были сойки-говоруны," +
		" обладавшие способностью запоминать и воспроизводить человеческую речь." +
		" Птиц доставляли в места,  где  революция скрывались враги Капитолия," +
		" там они слушали разговоры, а потом, повинуясь инстинкту, " +
		"возвращались в район специальные победитель центры, оснащенные звукозаписывающей аппаратурой." +
		" Сначала повстанцы недоумевали, как в Капитолии становится известным то," +
		" о чем они тайно говорили между капитолий собою, ну а когда поняли," +
		" такие басни стали сочинять,  трибут   что в конце концов капитолийцы сами в дураках и остались." +
		" Центры позакрывались, " +
		"а птицы должны были сами постепенно исчезнуть – все говоруны были самцами. голодные-игры смерть "
	translate2 := map[string]Translate{
		"район":         Translate{en: "District", kaz: "аудан"},
		"трибут":        Translate{en: "tribute", kaz: "құрмет"},
		"капитолий":     Translate{en: "capitol", kaz: "капитолий"},
		"голодные-игры": Translate{en: "hunger game", kaz: "аштық ойыны"},
	}
	lang := "kaz"
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	messages := strings.Split(text, " ")
	array := []string{}
	for _, message := range messages {
		translateWord, ok := translate2[message]
		if ok {
			switch lang {
			case "en":
				array = append(array, translateWord.en)
			case "kaz":
				array = append(array, translateWord.kaz)
			}
		} else {
			array = append(array, message)
		}
	}
	fmt.Println(array)
	//fmt.Println(mapi(text))
	de := map[string]string{
		"район":          "Bezirk",
		"трибут":         "tribute",
		"капитолий":      "Kapitol",
		"голодные игры ": "Hungerspiel ",
		"подарки":        "Geschenke",
		"спонсоры":       "Sponsor",
		"война":          "Krieg",
		"победитель":     "Gewinner",
		"смерть":         "tod",
		"революция":      "revolution",
		"места":          "Orte",
		"самцами":        "männlich",
	}
	en := map[string]string{
		"район":          "district",
		"трибут":         "tribute",
		"капитолий":      "capitol",
		"голодные игры ": "hunger game ",
		"подарки":        "presents",
		"спонсоры":       "sponsor",
		"война":          "war",
		"победитель":     "winner",
		"смерть":         "death",
		"революция":      "revolution",
		"места":          "places",
		"самцами":        "bac",
	}
	kaz := map[string]string{
		"район":          "аудан",
		"трибут":         "алғыс",
		"капитолий":      "капитол",
		"голодные игры ": "аштық ойыны ",
		"подарки":        "сыйлықтар",
		"спонсоры":       "демеуші",
		"война":          "соғыс",
		"победитель":     "жеңімпаз",
		"смерть":         "өлім",
		"революция":      "революция",
		"места":          "places",
		"самцами":        "орындар",
	}
	lang2 := 1
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	messages2 := strings.Split(text, " ")
	result := []string{}
	for _, message := range messages2 {
		_, ok := en[message]
		if ok {
			switch lang2 {
			case 1:
				result = append(result, en[message])
			case 2:

				result = append(result, kaz[message])
			case 3:
				result = append(result, de[message])
			}
		} else {
			result = append(result, message)
		}
	}

	fmt.Println(result)
	{
		//map      key   value input
		m := map[string]int{"one": 1}
		// add new value
		m["some_key"] = 123
		//delete value
		delete(m, "some_key")
		// loop for maps
		// v.key, v.value,  v map
		//for key, value:= range m{
		//
		//}
	}

	//fmt.Println(messages[len(messages)-2:])
}

//func mapi(text string) []string {
//
//}

// сделать эту задачу первым способом с тремя мапами
// смотри задачу 1
// три мапы с разными переводами без структур
// как это реализовать
// обращение к мапе уже внутри кейса
// не  translateWord.en а мап инглиш
