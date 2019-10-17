package piscine

func BasicAtoi(s string) int {

	var a int = 0

	for _, num := range s {

		var numb int = 0
		for i := '1'; i <= num; i++ {
			numb++
		}
		a = a*10 + numb

	}

	return a
}
