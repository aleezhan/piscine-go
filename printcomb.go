package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for l := '0'; l <= '9'; l++ {
				if i < j && j < l {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(l)
					z01.PrintRune(", ")
				}
			}
		}

	}
}
