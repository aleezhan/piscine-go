package piscine

func BasicAtoi(s string) int {

	stringM := []rune(s)
	stringJ := []rune()
	var numss string
	var i int = 0

	for _, word := range stringM {

		if word == '0' {
			continue
		} else {
			stringJ[i] = word
			i++
		}

	}

	for _, word := range stringJ {

		numss = word

	}

	return numss

}
