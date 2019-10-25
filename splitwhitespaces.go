package piscine

func SplitWhiteSpaces(str string) []string {
	ind := 1
	for _, word := range str {
		if word == ' ' {
			ind++
		}
	}
	var strr [ind]string
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
