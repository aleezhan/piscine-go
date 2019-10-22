package piscine

func Capitalize(s string) string {
	my_arr := []rune(s)
	if my_arr[0] >= 'a' && my_arr[0] <= 'z' {
		my_arr[0] = rune(my_arr[0] - 32)
	}
	k := 0
	for _, let := range s {
		let = let
		k++
	}
	for index, letter := range s {
		letter = letter
		if k-2 == index {
			continue
		}
		if (my_arr[index] < 'A' || my_arr[index] > 'Z') &&
			(my_arr[index+1] >= 'a' && my_arr[index+1] <= 'z') &&
			(my_arr[index] < 'a' || my_arr[index] > 'z') &&
			(my_arr[index] < '0' || my_arr[index] > '9') {
			my_arr[index+1] = rune(my_arr[index+1] - 32)
		}
		if (my_arr[index] >= 'A' && my_arr[index] <= 'Z') &&
			(my_arr[index+1] >= 'A' && my_arr[index+1] <= 'Z') {
			my_arr[index+1] = rune(my_arr[index+1] + 32)
		}
		if (my_arr[index] >= 'a' && my_arr[index] <= 'z') &&
			(my_arr[index+1] >= 'A' && my_arr[index+1] <= 'Z') {
			my_arr[index+1] = rune(my_arr[index+1] + 32)
		}
		if (my_arr[index] >= '0' && my_arr[index] <= '9') &&
			(my_arr[index+1] >= 'A' && my_arr[index+1] <= 'Z') {
			my_arr[index+1] = rune(my_arr[index+1] + 32)
		}

	}
	return string(my_arr)
}
