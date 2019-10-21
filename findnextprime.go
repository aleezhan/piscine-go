package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < nb+10; i++ {
		if i == 2 || i == 3 || i == 5 {
			a := i
			break
		}
		if i == 1 {
			a := i
			break
		}
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			a := i
			break
		}
	}
	return a
}
