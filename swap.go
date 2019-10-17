package priscine

func Swap(a *int, b *int) {

	var c *int

	*c = *a
	*a = *b
	*b = *c

}
