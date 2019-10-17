package piscine

func Swap(a *int, b *int) {

	var c int = *a
	var d int = *b
	*b = c
	*a = d

}
