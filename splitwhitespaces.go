package piscine

func SplitWhiteSpaces(str string) []string {
	var strr []string
	in := 0
	for _, words := range str {
		if words == ' ' {
			in++
			continue
		}
		strr[in] += string(words)
	}
	return strr

}
