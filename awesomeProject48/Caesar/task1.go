package main

import (
	"flag"
	"fmt"
	"strings"
)

// Цезарь это структура, содержит внутри алфавит, для использования де/энкрипта
// Caesar is a struct, holding alphabet being used, alongside with Encrypt and Decrypt functions
type Caesar struct {
	Alphabet []string
}

// EnglishAlphabet это модернизированый английский алфавит
// EnglishAlphabet is a modern English language alphabet

func (c *Caesar) doShiftedAlphabet(shift int) (shiftedAlphabet []string) {
	shiftedAlphabet = make([]string, len(c.Alphabet))
	if shift <= 0 {
		for shift < len(c.Alphabet) {
			shift = len(c.Alphabet) + shift
		}
	}
	for shift > len(c.Alphabet) {
		for shift > len(c.Alphabet) {
			shift = shift - len(c.Alphabet)
		}

	}
	for ak := range c.Alphabet {
		if (ak + shift) < len(c.Alphabet) {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift]
		} else {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift-len(c.Alphabet)]
		}
	}
	return
}
func (c *Caesar) findInAlphabet(alphabet []string, input string) int {
	for k, v := range alphabet {
		if v == input {
			return k
		}
	}
	return -1
}

func New(alphabet string) *Caesar {
	return &Caesar{Alphabet: strings.Split(alphabet, "")}
}
func main() {
	const EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var text = flag.String("text", "-", "Target value")
	var shifts = flag.Uint("shift", 0, "Target value")
	caesar := New(EnglishAlphabet)
	flag.Parse()
	fmt.Println("original text:", text)
	fmt.Println("shift :", shifts)
	caesar.Decrypt(*text, *shifts)
	fmt.Println(caesar.Decrypt(*text, *shifts))

	//fmt.Println("Encrypt text", Encrypt(text, shift))
	//fmt.Println("Decrypt text", Decrypt(text, shift))
}

//func NewEnglish() Caesar {
//	return New(EnglishAlphabet)
//}
