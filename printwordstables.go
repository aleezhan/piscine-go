package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	for i := 0; i < len(table); i++ {
		for _, word := range table[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}

}
