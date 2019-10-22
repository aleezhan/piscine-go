package piscine

func Index(s string, toFind string) int {
	index := 0
	arr := []rune(toFind)
	for in, letter := range s {
		if arr[0] == letter {
			index = in
			for i := 0; i < toFind.length(); i++ {
				if arr[i] == letter+i {
					if i == toFind.length() {
						return index
					}
				}
			}
		}
	}
	return -1
}
