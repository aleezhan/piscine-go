package piscine

func NRune(s string, n int) rune {
	l := '\x00'
	for index, letter := range s {
		if index+1 == n {
			l = letter
			break
		}
	}
	return l
}
