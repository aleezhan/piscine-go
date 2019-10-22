package piscine

import "github.com/01-edu/z01"

func TrimAtoi(s string) int {
	c := 0
	k := 0
	for _, num := range s {
		if c == 0 && num == '-' {
			c++
		}
		if num >= 48 && num <= 57 {
			k *= 10
			k += int(num)
			if c != 1 {
				c += 2
			}
		}
	}
	if c == 1 {
		k *= -1
	}
	return k
}
