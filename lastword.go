import (
	"os"

	"github.com/01-edu/z01"
)

func LASTWORD() {
	args := os.Args[1:]
	index1 := 0
	index2 := 0
	slice := ""
	if len(args) == 1 {
		for i := len(args[0]) - 1; i >= 0; i-- {
			if index1 == 0 && args[0][i] != ' ' {
				index1 = i + 1
			} else if index1 > 0 && args[0][i] == ' ' {
				index2 = i + 1
				break
			}
		}
		slice = args[0][index2:index1]
		for i := range slice {
			z01.PrintRune(rune(slice[i]))
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}

}
