package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	index--
	if index == 1 {
		return 1
	}
	return index + Fibonacci(index)
}
