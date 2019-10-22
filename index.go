package piscine

func Index(s string, toFind string) int {
	index := 0
	arr := []rune(toFind)
	arr2 := []rune(s)
	for in, letter := range s {
		if arr[0] == letter {
			index = in
			for i := 0; arr[i] != '\x00'; i++ {
				if arr[i] == arr2[in+i] {
					if i+1 == '\x00' {
						return index
					}
				}
			}
		}
	}
	return -1
}
