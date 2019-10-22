package piscine

func AlphaCount(str string) int {
	counter := 0
	for index, letter := range str {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 172) {
			counter++
		}
	}
	return counter
}
