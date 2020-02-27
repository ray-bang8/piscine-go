package main

import (
	"os"

	"github.com/01-edu/z01"
)

func REVWSTR() {
	if len(os.Args) == 2 {
		word := ""
		str := os.Args[1]
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] == ' ' && word != "" {
				for _, v := range word {
					z01.PrintRune(v)

				}
				word = ""
				if i != 0 {
					z01.PrintRune(32)
				}
			} else if str[i] != ' ' {
				word = string(str[i]) + word

			}

		}
		if word != "" {
			for _, v := range word {
				z01.PrintRune(v)

			}
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
______________________________
func REVWSTR2() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	} else {
		s := os.Args[1]
		res := ""
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == ' ' {
				res += s[i+1:] + " "
				s = s[:i]
			}
		}
		res += s
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
