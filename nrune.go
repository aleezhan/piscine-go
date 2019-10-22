package piscine

func NRune(s string, n int) rune {
	for index, letter := range s {
		if index+1 == n {
			return letter
		}
	}
}
