package piscine

func StrRev(s string) string {

	i := 0
	var c rune

	stringT := []rune(s)

	for _, word := range stringT {

		word = word

		i++

	}

	for j := i; j > 1; j-- {
		for k := 0; k < j-1; k++ {
			c = stringT[k]
			stringT[k] = stringT[k+1]
			stringT[k+1] = c

		}
	}

	s = ""

	for _, word := range stringT {

		word = word

		s = s + string(word)

	}

	return s

}
