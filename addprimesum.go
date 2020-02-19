import (
	"fmt"
	"os"
)

func ADDPRIMESUM() {

	if len(os.Args) == 2 {
		nbr := Atoi(os.Args[1])
		printPrime(nbr)

	} else {
		fmt.Println("0")
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

func printPrime(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i

			//fmt.Println(sum)
		}

	}
	fmt.Println(sum)
	return sum
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