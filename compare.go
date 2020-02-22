package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune('\n')
		return
	}
	a := os.Args[1]
	b := os.Args[2]
	if a < b {
		z01.PrintRune('-')
		z01.PrintRune('1')
	} else if a > b {
		z01.PrintRune('1')
	} else if a == b {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
