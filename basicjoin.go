package piscine

func BasicJoin(strs []string) string {
	var words string
	for _, word := range strs {
		words += word
	}
	return words
}
