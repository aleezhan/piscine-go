package piscine

func IsLower(str string) bool {
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			continue
		}
		return false
	}
	return true
}
