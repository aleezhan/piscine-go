package piscine

import (
	piscine ".."
)

func SplitWhiteSpaces(str string) []string {
	ind := 1
	c := 0
	for _, word := range str {
		if (word == ' ' || word == '\n' || word == '\t') && c == 0 {
			ind++
			c++
		} else if word != ' ' && word != ' ' && word != '\t' {
			c = 0
		}
	}
	strr := make([]string, ind)
	in := 0
	b := 0
	for index, words := range str {
		if (words == ' ' || words == '\n' || words == '\t') && b == 0 && index != 0 {
			in++
			b++
		} else if words != ' ' && words != '\n' && words != '\t' {
			strr[in] += string(words)
			b = 0
		}
	}
	PrintWordsTbales(strr)

}
