

import (
	"os"

	"github.com/01-edu/z01"
)

func UNION() {

	if len(os.Args[1:]) == 2 {
		fin := os.Args[1] + os.Args[2]
		//fmt.Println(fin)
		temp := []rune{}
		for _, k := range fin {
			if isUNique(k, temp) {
				temp = append(temp, k)
			}
		}
		for _, v := range temp {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}

func isUNique(r rune, s []rune) bool {
	for _, k := range s {
		if k == r {
			return false
		}
	}
	return true
}
