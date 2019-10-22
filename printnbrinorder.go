package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	a := 0
	b := 0
	if n > 0 {
		for {
			a /= 10
			b++
			if a <= 0 {
				break
			}
		}
	}
	for i := 1; i < 10; i++ {
		c := n
		d := c
		for j := 0; j < b; j++ {
			if d%10 == i {
				d = c % 10
				z01.PrintRune(rune(d + 48))
				c /= 10
			}
		}
	}
}
