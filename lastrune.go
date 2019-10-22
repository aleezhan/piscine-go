package piscine

func LastRune(s string) rune {
	l := 'd'
	for _, letter := range s {
		l = letter
	}
	return l
}
