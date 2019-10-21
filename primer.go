package main

import "github.com/01-edu/z01"

func PrintStr(str string) {

	var stringT rune = []byte(str)

	for index, word := range stringT {

		z01.PrintRune('%v', word)

	}

}
