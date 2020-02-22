
import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func Ispower022() {
	if len(os.Args) == 2 {
		x, _ := strconv.Atoi(os.Args[1])
		prTrue := "true"
		prFalse := "false"

		for x%2 == 0 {
			if x == 0 {
				break
			}
			x = x / 2
		}
		if x == 1 {
			for i := range prTrue {
				z01.PrintRune(rune(prTrue[i]))
			}
		}
		if x == 0 || x != 1 {
			for i := range prFalse {
				z01.PrintRune(rune(prFalse[i]))
			}
		}
		z01.PrintRune(10)

	} else {
		z01.PrintRune(10)
	}
}

func ispowerof2() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		x, err := strconv.Atoi(arg[0])
		if err != nil {
			z01.PrintRune(10)
			return
		} else if x != 0 && (x&(x-1)) == 0 {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
		} else {
			z01.PrintRune('f')
			z01.PrintRune('a')
			z01.PrintRune('l')
			z01.PrintRune('s')
			z01.PrintRune('e')
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}

}
