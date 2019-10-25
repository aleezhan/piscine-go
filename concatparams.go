package piscine

func ConcatParams(args []string) string {

	str := ""
	size := 0
	for index := range args {
		size = index
	}
	for i := 0; i <= size; i++ {
		if i == size {
			str += args[i]
			break
		}
		str += args[i] + "\n"

	}
	return str

}
