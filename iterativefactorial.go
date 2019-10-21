package piscine

import (
	"fmt"
)

func IterativeFactorial(nb int) int {

	f := 0
	if nb > 0 {

		for i := 0; i <= nb; i++ {

			f = f + i

		}

		print(f)

	} else {
		return 0
	}
}
