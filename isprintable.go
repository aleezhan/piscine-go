package piscine

func IsPrintable(str string) bool {
	for _, letter := range str {
		if letter >= 0 && letter <= 32 {
			return false
		}
	}
	return true
}
