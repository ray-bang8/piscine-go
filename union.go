import (
	"fmt"
	"os"
)

func UNION() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		sum := s1 + s2
		mas := []rune(sum)
		for i := 0; i < len(mas); i++ {
			for k := i + 1; k < len(mas); k++ {
				if mas[i] == mas[k] {
					mas[k] = 0
				}
			}
		}
		fmt.Println(string(mas))

	} else {
		fmt.Println()
	}
}