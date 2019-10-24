package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	i := 0
	for k := range arg {
		i = k
	}
	for j := i; j >= 1; j-- {
		for _, b := range arg[j] {
			z01.PrintRune(b)

		}
		z01.PrintRune('\n')
	}
}
