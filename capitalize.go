package piscine

func Capitalize(s string) string {
	my_arr := []rune(s)
	if my_arr[0] >= 'a' && my_arr[0] <= 'z' {
		my_arr[0] = rune(my_arr[0] - 32)
	}

	k := 0

	for index := range my_arr {
		k = index
	}

	for i := 1; i <= k; i++ {
		if (my_arr[i-1] < 'A' || my_arr[i-1] > 'Z') &&
			(my_arr[i] >= 'a' && my_arr[i] <= 'z') &&
			(my_arr[i-1] < 'a' || my_arr[i-1] > 'z') &&
			(my_arr[i-1] < '0' || my_arr[i-1] > '9') {
			my_arr[i] = rune(my_arr[i] - 32)
		}
		if (my_arr[i-1] >= 'A' && my_arr[i-1] <= 'Z') &&
			(my_arr[i] >= 'A' && my_arr[i] <= 'Z') {
			my_arr[i] = rune(my_arr[i] + 32)
		}
		if (my_arr[i-1] >= 'a' && my_arr[i-1] <= 'z') &&
			(my_arr[i] >= 'A' && my_arr[i] <= 'Z') {
			my_arr[i] = rune(my_arr[i] + 32)
		}
		if (my_arr[i-1] >= '0' && my_arr[i-1] <= '9') &&
			(my_arr[i] >= 'A' && my_arr[i] <= 'Z') {
			my_arr[i] = rune(my_arr[i] + 32)
		}

	}

	return string(my_arr)
}
