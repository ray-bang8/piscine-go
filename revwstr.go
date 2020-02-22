
import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func REVWSTR() {

	if len(os.Args) == 2 {
		str := SplitWhiteSpaces(os.Args[1])
		for i := len(str) - 1; i >= 0; i-- {
			// fmt.Print(string(str[i]) + " ")
			for j := range str[i] {
				z01.PrintRune(rune(str[i][j]))

			}
			if i != 0 {
				z01.PrintRune(32)
			}

		}
		z01.PrintRune(10)

	} else {
		fmt.Println()
	}
}
func SplitWhiteSpaces(str string) []string {
	prev := ' '
	count := 0
	for _, v := range str {
		if v != ' ' && prev == ' ' {
			count++
		}
		prev = v
	}
	ar := make([]string, count)
	i := 0
	word := ""
	for _, v := range str {
		if v == ' ' && word != "" {
			ar[i] = word
			word = ""
			if count != 1 {
				i++
			}
		} else if v != ' ' {
			word += string(v)
		}
	}
	if word != "" {
		ar[i] = word
	}
	return ar
}
