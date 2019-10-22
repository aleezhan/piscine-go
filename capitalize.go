package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	l := 0
	l = l
	if arr[0] >= 97 && arr[0] <= 122 {
		l = 0
	} else {
		l = 1
	}
	for index, letter := range s {
		if letter >= 97 && letter <= 122 && l == 0 {
			arr[index] = letter - 32
			l++
		}
		if letter < 97 || letter > 122 && l == 1 {
			l--
		}

	}
	return string(arr)
}
