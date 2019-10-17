package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for l := '0'; l <= '9'; l++ {
				for m := '0'; m <= '9'; m++ {
					if i == '9' && j == '8' && l == '9' && m == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(040)
						z01.PrintRune(l)
						z01.PrintRune(m)
						z01.PrintRune(10)
						break
					}
					if i == l {
						if j < m {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(040)
							z01.PrintRune(l)
							z01.PrintRune(m)
							z01.PrintRune(',')
							z01.PrintRune(040)
						}
					}
					if i < l {

						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(040)
						z01.PrintRune(l)
						z01.PrintRune(m)
						z01.PrintRune(',')
						z01.PrintRune(040)

					}
				}
			}
		}
	}
}
