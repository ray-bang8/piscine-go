

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func PRINTCHESSBARD() {

	arg := os.Args[1:]
	er := "Error"
	if len(arg) == 2 {

		col, err1 := strconv.Atoi(arg[0])
		if err1 != nil {
			for _, v := range er {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
		}
		row, err2 := strconv.Atoi(arg[1])
		if err2 != nil {
			for _, v := range er {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
		}
		if col == 0 && row == 0 {
			for _, v := range er {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
			return
		}
		for i := 1; i <= row; i++ {
			for j := 1; j <= col; j++ {
				if i%2 != 0 {
					if j%2 != 0 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
				} else {
					if j%2 != 0 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('#')
					}
				}

			}
			z01.PrintRune('\n')
		}

	} else {
		for _, v := range er {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}

}
