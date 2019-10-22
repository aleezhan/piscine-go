package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for index, letter := range s {
		if letter >= 65 && letter <= 90 {
			arr[index] = letter + 32
		}
	}
	return string(arr)
}
