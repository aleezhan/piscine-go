package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var k = 0
	var j = 0
	var num [20]int
	for i := n; i >= 0; i / 10 {
		k = i
		k = i % 10
		j = 0
		num[j] = k
		j = j + 1
	}
	var c = 0
	for a := j - 1; a > 0; a-- {
		for b := 0; b < a; b++ {
			c = num[b]
			num[b] = num[b+1]
			num[b+1] = c
		}
	}

	if n < 0 {
		z01.PrintRune(45)
	}

	for g := 0; g < j; g++ {
		if num[g] == 0 {
			z01.PrintRune(48)
		} else if num[g] == 1 {
			z01.PrintRune(49)
		} else if num[g] == 2 {
			z01.PrintRune(50)
		} else if num[g] == 3 {
			z01.PrintRune(51)
		} else if num[g] == 4 {
			z01.PrintRune(52)
		} else if num[g] == 5 {
			z01.PrintRune(53)
		} else if num[g] == 6 {
			z01.PrintRune(54)
		} else if num[g] == 7 {
			z01.PrintRune(55)
		} else if num[g] == 8 {
			z01.PrintRune(56)
		} else if num[g] == 9 {
			z01.PrintRune(57)
		}

	}

}
