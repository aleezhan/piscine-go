package piscine

func SplitWhiteSpaces(str string) []string {
	ind := 1
	for _, word := range str {
		if word == ' ' || word == '\n' || word == '\t' {
			ind++
		}
	}
	strr := make([]string, ind)
	in := 0
	b := 0
	for _, words := range str {
		if words == ' ' || word == '\n' || word == '\t' && b == 0 {
			in++
			b++
			continue
		}
		strr[in] += string(words)
		b = 0
	}
	return strr

}
