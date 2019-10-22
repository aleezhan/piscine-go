package piscine

func AlphaCount(str string) int {
	counter := 0
	for _, letter := range str {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 172) {
			counter++
		}
	}
	return counter
}
