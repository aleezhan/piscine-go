package piscine

func Sqrt(nb int) int {
	for i := 0; i < 1000; i++ {
		if i*i == nb {
			return i
		} else {
			return 0
		}
	}
}
