
import (
	"os"

	"github.com/01-edu/z01"
)

func ALPHAMiror() {
	if len(os.Args[1:]) == 1 {
		input := os.Args[1]
		str := []rune(input)
		for i := range str {
			if !(str[i] >= 'a' && str[i] <= 'z') && !(str[i] >= 'A' && str[i] <= 'Z') {
				continue
			}
			if str[i] >= 'a' && str[i] <= 'm' {
				count := 25
				for j := 'a'; j < str[i]; j++ {
					count -= 2
				}
				str[i] = str[i] + rune(count)
			} else if str[i] >= 'n' && str[i] <= 'z' {
				count := 25
				for j := 'z'; j > str[i]; j-- {
					count -= 2
				}
				str[i] = str[i] - rune(count)
			} else if str[i] >= 'A' && str[i] <= 'M' {
				count := 25
				for j := 'A'; j < str[i]; j++ {
					count -= 2
				}
				str[i] = str[i] + rune(count)
			} else if str[i] >= 'N' && str[i] <= 'Z' {
				count := 25
				for j := 'Z'; j > str[i]; j-- {
					count -= 2
				}
				str[i] = str[i] - rune(count)
			}

		}
		for i := range str {
			z01.PrintRune(str[i])
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}

}
