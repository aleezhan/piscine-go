package piscine

func IterativeFactorial(nb int) int {
	f := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 25 {
		for i := 1; i <= nb; i++ {
			f = f * i
		}
	} else {
		return 0
	}

	return f
}