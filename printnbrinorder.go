package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	a := 0
	b := 0
	a = n
	if n > 0 {
		for {
			if a <= 0 {
				break
			}
			a = a / 10
			b++
		}
	} else {
		z01.PrintRune(rune(0 + 48))
	}

	for i := 0; i < 10; i++ {
		c := n
		for j := 0; j < b; j++ {
			if c%10 == i {
				z01.PrintRune(rune((c % 10) + 48))
				c /= 10
			} else {
				c /= 10
			}
		}
	}
}
