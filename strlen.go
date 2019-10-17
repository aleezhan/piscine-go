package piscine

func StrLen(str string) int {

	i := 0

	stringT := []rune(str)

	for _, word := range stringT {

		i++

	}

	return i

}
