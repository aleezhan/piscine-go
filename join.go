package piscine

func Join(strs []string, sep string) string {
	var words string
	k := len(strs)
	for index, word := range strs {
		if index == strs-1 {
			words += word
			continue
		}
		words += word + sep
	}
	return words
}
