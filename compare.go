package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	ar := []rune(a)
	arr := []rune(b)
	indexx := 0
	b := 0
	for _, lett := range b {
		lett = lett
		indexx++
	}
	for i := 0; i < indexx; i++ {
		if ar[0] == arr[0] {
			continue
		}
		b++
		break
	}
	if b == 0 {
		return 1
	}
	return -1
}
