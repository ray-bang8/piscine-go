package main

import (
	"os"

	"github.com/01-edu/z01"
)

func UniqueOccurences(s string) bool {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}

	// fmt.Println(m)

	prev := 0
	for _, n := range m {
		if prev != 0 && prev == n {
			return false
		}

		prev = n
	}

	return true
}
func main() {
	args := os.Args[1:]
	tr := "true"
	fs := "false"
	if len(args) == 1 {
		if UniqueOccurences(args[0]) == true {
			for _, v := range tr {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
		} else {
			for _, v := range fs {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
		}
	} else {
		z01.PrintRune(10)
	}
}
