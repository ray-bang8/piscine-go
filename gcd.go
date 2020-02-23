
import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func GCD() {
	if len(os.Args) == 3 {

		s1, _ := strconv.Atoi(os.Args[1])
		s2, _ := strconv.Atoi(os.Args[2])
		fmt.Println(gcd(s1, s2))

	} else {
		z01.PrintRune('\n')
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
------------------------------------
func GCD2() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune('\n')
		return
	} else {
		min, _ := strconv.Atoi(os.Args[1])
		max, _ := strconv.Atoi(os.Args[2])
		if max < min {
			min, max = max, min
		}
		gcd := 0
		for i := min; i > 0; i-- {
			if min%i == 0 && max%i == 0 {
				gcd = i
				break
			}
		}
		fmt.Println(gcd)
	}
}
_____________________
package gcd

func gcd(x, y int) int{
  for y != 0 {
    x, y = y, x%y
  }
  return x
}