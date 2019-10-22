package piscine

func ToUpper(s string) string {
	for index, letter := range s {
		if letter >= 97 && letter <= 122 {
			letter -= 32
		}
	}
}
