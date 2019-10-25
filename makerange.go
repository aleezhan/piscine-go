package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	arr := make([]int, size)
	for i := 0; i < max; i++ {
		arr[i] = min
		min++
	}
	return arr

}
