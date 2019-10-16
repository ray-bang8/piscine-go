package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for b := '0'; b <= '9'; b++ {
					if k > i {
						PrintNumbers(i, j, k, b)
					} else if k == i && b > j {
						PrintNumbers(i, j, k, b)
					}
				}

			}
		}

	}

}
func PrintNumbers(i, j, k, b rune) {
	z01.PrintRune(i)
	z01.PrintRune(j)
	z01.PrintRune(' ')
	z01.PrintRune(k)
	z01.PrintRune(b)
	if i == '9' && j == '8' && k == '9' && b == '9' {
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
