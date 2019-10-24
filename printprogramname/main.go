package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	cox := os.Args
	for _, i := range cox[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

}
