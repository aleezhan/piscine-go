package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arg := os.Args
	arg2 := os.Args
	l := 0
	for _, letter := range arg {
		letter = letter
		l++
	}
	for i := 1; i < l; i++ {
		for j := i; j < l; j++ {
			if arg2[i] > arg[j] {
				arg2[j], arg[i] = arg[i], arg2[j]
			}
		}
	}
	for k := 1; k < l; k++ {
		for _, la := range arg2[k] {
			z01.PrintRune(la)
		}
		z01.PrintRune('\n')
	}

}
