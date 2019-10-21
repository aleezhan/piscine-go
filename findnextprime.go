package piscine

func FindNextPrime(nb int) int {
	a := 0
	if nb < 0 {
		nb = 0
	}
	for i := nb; i < nb+1000; i++ {
		if i == 0 {
			continue
		}
		if i == 1 {
			a = i + 1
			break
		}
		if i == 2 || i == 3 || i == 5 || i == 7 {
			a = i
			break
		}
		if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 && i > 0 {
			a = i
			break
		}
	}
	if a == 0 {
		return a
	}
	return a
}
