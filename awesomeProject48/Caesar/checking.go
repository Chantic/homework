package main

import "strings"

func (c *Caesar) Decrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabet(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(shiftedAlphabet, v)
		if position != -1 {
			ret = ret + c.Alphabet[position]
		} else {
			ret = ret + v
		}
	}
	return
}
func (c *Caesar) Encrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabet(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(c.Alphabet, v)
		if position != -1 {
			ret = ret + shiftedAlphabet[position]
		} else {
			ret = ret + v
		}

	}
	return
}
