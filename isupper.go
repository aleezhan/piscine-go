package piscine

func IsUpper(str string) bool {
	for _, letter := range str {
		if letter >= 'A' && letter <= 'Z' {
			continue
		}
		return false
	}
	return true
}
