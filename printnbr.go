package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	c := '0'
	if n == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= n%10; i++ {
		c++
	}
	for i := -1; i >= n%10; i-- {
		c++
	}
	if n/10 != 0 {
		checkit(n / 10)
	}
	z01.PrintRune(c)
}
