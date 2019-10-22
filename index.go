package piscine

func Index(s string, toFind string) int {
	index := 0
	k := 0
	arr2 := []rune(toFind)
	arr := []rune(s)
	for _, let := range s {
		let = let
		index++
	}
	for _, lett := range toFind {
		lett = lett
		k++
	}
	a := 1
	if index >= k {
		for i := 0; i < index; i++ {
			if arr[i] == arr2[0] {
				a = 1
				for j := 0; j < k; j++ {
					if arr2[j] == arr[i+j] {
						continue
					}
					a = 0
					break
				}
				if a == 1 {
					return i
				}
			}
		}
	}
	return -1
}
