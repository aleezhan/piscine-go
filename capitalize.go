package piscine

func Capitalize(s string) string {
	my_arr := []rune(s)
	for index, letter := range s {
		letter = letter
		if (my_arr[index-1] < 'A' || my_arr[index-1] > 'Z') &&
			(my_arr[index] >= 'a' && my_arr[index] <= 'z') &&
			(my_arr[index-1] < 'a' || my_arr[index-1] > 'z') &&
			(my_arr[index-1] < '0' || my_arr[index-1] > '9') {
			my_arr[index] = rune(my_arr[index] - 32)
		}
		if (my_arr[index-1] >= 'A' && my_arr[index-1] <= 'Z') &&
			(my_arr[index] >= 'A' && my_arr[index] <= 'Z') {
			my_arr[index] = rune(my_arr[index] + 32)
		}
		if (my_arr[index-1] >= 'a' && my_arr[index-1] <= 'z') &&
			(my_arr[index] >= 'A' && my_arr[index] <= 'Z') {
			my_arr[index] = rune(my_arr[index] + 32)
		}
		if (my_arr[index-1] >= '0' && my_arr[index-1] <= '9') &&
			(my_arr[index] >= 'A' && my_arr[index] <= 'Z') {
			my_arr[index] = rune(my_arr[index] + 32)
		}

	}
	return string(my_arr)
}
