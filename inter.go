

import (
	"os"

	"github.com/01-edu/z01"
)

func INTER() {
	answer := ""
	if len(os.Args[1:]) == 2 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		for i := 0; i < len(str1); i++ {
			for j := 0; j < len(str2); j++ {
				if str2[j] == str1[i] {
					doesNotExist := true
					for k := range answer {
						if str1[i] == answer[k] {
							doesNotExist = false // str1[k] = 0
						}
					}
					if doesNotExist {
						answer += string(str1[i])
					}
				}

			}
		}
		for i := range answer {
			z01.PrintRune(rune(answer[i]))
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
