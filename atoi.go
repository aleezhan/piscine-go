package piscine

func Atoi(s string) int {

	var a int = 0
	var n int = 0

	for _, num := range s {
		var numb int = 0
		if s[0] == '-' && n == 0 {
			n++
		} else if num > '9' || num < '0' {
			a = 0
			break
		}

		for i := '1'; i <= num; i++ {
			numb++
		}
		a = a*10 + numb

	}
	if n == 1 {
		a = a * -1
	}

	return a
}
