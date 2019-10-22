package piscine

func Join(strs []string, sep string) string {
	var words string
	for in, word := range strs {
		k := in
	}
	for index, word := range strs {
		if index == in {
			words += word
			continue
		}
		words += word + sep
	}
	return words
}
