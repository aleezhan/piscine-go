package piscine

import "github.com/01-edu/z01"

func IterativeFactorial(nb int) int {

	f := 0
	if nb > 0 {

		for i := 0; i <= nb; i++ {

			f = f + i

		}

		z01.PrintRune(char(f))

	} else {
		return 0
	}
}
