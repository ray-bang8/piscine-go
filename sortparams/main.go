package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	a := 0
	for range arg {
		a++
	}
	newstr := []string(arg)
	for i := 0; i < a; i++ {
		for j := i; j < a; j++ {
			if newstr[i] > newstr[j] {
				swap(newstr, i, j)

			}
		}
	}
	for v := range arg {
		for _, c := range newstr[v] {
			z01.PrintRune(c)
		}
		z01.PrintRune(10)
	}
}

func swap(n []string, i, j int) {
	tmp := n[i]
	n[i] = n[j]
	n[j] = tmp
}
