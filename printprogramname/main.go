package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	z01.PrintRune(arg[0])
	z01.PrintRune('\n')

}
