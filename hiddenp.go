import (
	"fmt"
	"os"
)

func HIDDENP() {
	if len(os.Args) == 3 {
		answer := ""
		s1 := []rune(os.Args[1])
		s2 := []rune(os.Args[2])
		j := 0

		for i := 0; i < len(s1); i++ {
			for ; j < len(s2); j++ {
				if s1[i] == s2[j] {
					answer += string(s1[i])
					break
				}
			}
		}
		fmt.Println(answer)
		if answer == string(s1) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	} else {
		fmt.Println()
	}
}