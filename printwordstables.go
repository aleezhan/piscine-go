package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	l := 0
	for _, words := range table {
		words = words
		l++
	}
	for i := 0; i < l; i++ {
		for _, word := range table[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}

}
