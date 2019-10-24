package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	cox := os.Args
	for index, cox := range cox {
		if index != 0 {
			for _, i := range cox {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
		}
	}
}
