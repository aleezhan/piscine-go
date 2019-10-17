package piscine

func SortIntegerTable(table []int) {

	var a int = len(table)
	var b int

	for i := a - 1; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			b = table[j]
			table[j] = table[j+1]
			table[j+1] = b
		}
	}

	return table

}
