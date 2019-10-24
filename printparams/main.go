package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arg := os.Args
	for i := 0; i < len(arg); i++ {
		for _, word := range arg[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}

}
