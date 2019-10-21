package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 200000; i++ {
		if i == 2 || i == 3 || i == 5 {
			return i
		}
		if i == 1 {
			return false
		}
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			return i
		}
	}
}
