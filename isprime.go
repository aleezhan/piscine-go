package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return true
	}
	if nb%2 != 0 && nb%3 != 0 && nb%5 != 0 {
		return true
	}
	return false
}
