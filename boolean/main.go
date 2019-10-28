package main

import (
	"os"
	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func even(nbr int) boolean {
	if nbr == 8 {
		return true
	}else{
		return false
	}
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	arg := os.Args
	lengthOfArg := 0
	var length int
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	for index := range arg {
		length := index
	} 
	for i := 1; i <= length; i++ {
		for _, letter := range arg[i] {
			lengthOfArg += 1
		}
	}
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
