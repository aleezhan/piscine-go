package piscine

func FindNextPrime(nb int) int {
	a := 0
	for i := nb; i < nb+1000; i++ {
		if i == 2 || i == 3 || i == 5 {
			a = i
			break
		}
		if i == 1 {
			a = i
			break
		}
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			a = i
			break
		}
	}
	if a == 0 {
		return a
	}
	return a
}
