package piscine

func BasicAtoi2(s string) int {

	var a int = 0

	for _, num := range s {
		var numb int = 0
		if num > '9' || num < '0' {
			a = 0
			break
		}

		for i := '1'; i <= num; i++ {
			numb++
		}
		a = a*10 + numb

	}

	return a
}
