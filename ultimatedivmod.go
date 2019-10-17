package piscine

func UltimateDivMod(a *int, b *int) {

	var c int = *a
	var d int = *b
	*a = c / d
	*b = c % d

}
