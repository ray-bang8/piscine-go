package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		answer := ""
		s1 := []rune(os.Args[1])
		s2 := []rune(os.Args[2])
		j := 0

		for i := 0; i < len(s1); i++ {

			for ; j < len(s2); j++ {
				if s1[i] == s2[j] {
					answer += string(s1[i])
					j++
					break
				}
			}
		}
		if answer == string(s1) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
