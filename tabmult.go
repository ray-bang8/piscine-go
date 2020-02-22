
import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func tabmult() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	} else {
		a, _ := strconv.Atoi(os.Args[1])
		b := os.Args[1]
		if a < 0 {
			z01.PrintRune('\n')
			return
		}
		for i := 1; i <= 9; i++ {
			res := i * a
			z01.PrintRune(rune(i + 48))
			z01.PrintRune(' ')
			z01.PrintRune('x')
			z01.PrintRune(' ')
			for _, v := range b {
				z01.PrintRune(v)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			str := it(res)
			for _, v := range str {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')

		}
	}
}
func Atoi(s string) int {
	len := 0
	for range s {
		len++
	}
	plus := 0
	minus := 0
	final := 0
	count := 0
	str := []rune(s)
	for i := 0; i < len; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			count = 0
			for j := '0'; j < str[i]; j++ {
				count++
			}
			final = final*10 + count
		} else if str[i] == '-' && str[i] == 0 {
			minus++
		} else if str[i] == '+' && str[i] == 0 {
			plus++
		} else {
			return 0
		}

	}
	if minus == 1 {
		final = -final
	}
	return final
}
func it(n int) string {
	var ost int
	s := ""
	if n == 0 {
		return "0"
	}
	min := 0
	for {
		if n == 0 {
			break
		}
		if n < 0 {
			ost = -(n % 10)
		} else {
			ost = n % 10
		}
		s = string(ost+48) + s
		if n/10 == 0 && n < 0 {
			min++
		}
		n = n / 10
	}
	if min == 1 {
		s = "-" + s
	}
	return s
}
