package piscine

func IsAlpha(str string) bool {
	for _, letter := range str {
		if (letter >= 'a' && letter <= 'z') ||
			(letter >= 'A' && letter <= 'Z') ||
			(letter >= '0' && letter <= '9') {
			continue
		}
		return false
	}
	return true
}
