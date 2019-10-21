package piscine

func IterativeFactorial(nb int) int {

	f := 0
	if nb > 0 {

		for i := 0; i <= nb; i++ {

			f = f + i

		}
		return f

	} else {
		return 0
	}
}
