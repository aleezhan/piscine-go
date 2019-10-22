package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for index, letter := range s {
		if letter >= 97 && letter <= 122 {
			arr[index] = letter - 32
		}
	}
	return string(arr)
}
