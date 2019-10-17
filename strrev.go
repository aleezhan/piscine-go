package piscine

func StrRev(s string) string {

	stringM := []rune(s)

	i := 0
	var c rune

	stringT := []rune(s)

	for _, word := range stringT {

		word = word

		i++

	}

	for j := i; j > 1; j-- {
		for k := 0; k < j-1; k++ {
			c = str[k]
			str[k] = str[k+1]
			str[k+1] = c

		}
	}

	s = ""

	for _, word := range stringT {

		word = word

		s = s + word

	}

	return s

}
