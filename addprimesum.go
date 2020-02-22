

import (
	"os"

	"github.com/01-edu/z01"
)

func ADDPRIMESUM() {

	if len(os.Args) == 2 {
		nbr := Atoi(os.Args[1])
		printPrime(nbr)

	} else {
		z01.PrintRune('0')
		z01.PrintRune(10)
	}
}

func IsPrime(value int) bool {
	if value <= 1 {
		return false
	}
	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func printPrime(n int) {
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i

			//fmt.Println(sum)
		}

	}
	res := itoa(sum)
	for i := range res {
		z01.PrintRune(rune(res[i]))
	}
	z01.PrintRune(10)

}
func Atoi(s string) int {
	min := 0
	plu := 0
	res := 0
	for i, v := range s {
		if v >= '0' && v <= '9' {
			num := 0
			for g := '0'; g < v; g++ {
				num++
			}
			res = res*10 + num
		} else if v == '-' && i == 0 {
			min++
		} else if v == '+' && i == 0 {
			plu++
		} else {
			return 0
		}
	}
	if min == 1 {
		res = -res
	}
	return res
}
func itoa(n int) string {
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
