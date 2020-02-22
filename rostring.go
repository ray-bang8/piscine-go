
import (
	"os"

	"github.com/01-edu/z01"
)

func ROSTRING() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	} else {
		words := splitWhiteSpaces(os.Args[1])

		//moving all words to 1 position to the left
		newWords := make([]string, len(words))
		j := 0
		for i := 1; i < len(words); i++ {
			newWords[j] = words[i]
			j++
		}
		newWords[j] = words[0]

		//printing the result
		res := newWords[0]
		for i := 1; i < len(words); i++ {
			res = res + " " + newWords[i]
		}

		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func splitWhiteSpaces(s string) []string {
	prev := ' '
	count := 0
	for _, v := range s {
		if v != ' ' && prev == ' ' {
			count++
		}
		prev = v
	}
	ar := make([]string, count)
	i := 0
	word := ""
	for _, v := range s {
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
