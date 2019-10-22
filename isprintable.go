package piscine

func IsPrintable(str string) bool {
	for _, letter := range str {
		if letter != '\\' {
			continue
		}
		return false
	}
	return true
}
