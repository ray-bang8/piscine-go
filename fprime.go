import (
	"fmt"
	"os"
)

func FPRIME() {
	if len(os.Args) == 2 {
		v := Atoi(os.Args[1])
		count := 0
		l := append(PF(v))
		for i := 0; i < len(l); i++ {
			count++
			if count >= i && i > 0 {
				fmt.Print("*")
			}
			fmt.Print(l[i])
		}
		fmt.Println()
	} else {
		fmt.Println()
	}
}
func PF(x int) []rune {
	f := []rune{}
	c := 2
	for c <= x {
		if x%c == 0 {
			f = append(f, rune(c))
			x = x / c
		} else {
			c++
		}

	}
	return f
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