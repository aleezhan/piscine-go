package piscine 

import "github.com/01-edu/z01"

func Raid1a(x, y int) {

	for i := 0; i < y; i++ {

		for j := 0; j < x; j++ {

			if (i == 0 || i == (y - 1)) && (j == 0 || (j == (x - 1))) {
				z01.PrintRune('o')
			} else if (i == 0 || i == (y - 1)) && (j != 0 || (j != (x - 1))) {
				z01.PrintRune('-')
			} else if (i == 0 || i == (y - 1)) && (j == 0 || (j == (x - 1))) {
				z01.PrintRune('|')
			}

		}

	}

}