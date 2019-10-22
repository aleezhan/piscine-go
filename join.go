package piscine

func Join(strs []string, sep string) string {
	var words string
	k := 0
	for in, word := range strs {
		word = word
		k = in
	}
	for index, word := range strs {
		if index == k {
			words += word
			continue
		}
		words += word + sep
	}
	return words
}
