package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arg := os.Args
	r := 0
	for _, l := range arg {
		l = l
		r++
	}
	for i := 1; i < r; i++ {
		for _, word := range arg[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}

}
