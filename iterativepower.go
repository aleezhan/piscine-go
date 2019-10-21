package piscine

func IterativePower(nb int, power int) int {
	b := nb
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	for i := 2; i <= power; i++ {
		nb = nb * b
	}
	return nb

}
