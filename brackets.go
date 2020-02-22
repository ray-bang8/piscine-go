
import (
	"fmt"
	"os"
)

func BRACKETS() {
	if len(os.Args) == 1 {
		fmt.Println()
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arr := []rune(os.Args[i])
		var stack []rune

		if os.Args[i] == "" {

			fmt.Println("OK")
		} else {
			for j := 0; j < len(os.Args[i]); j++ {
				if arr[j] == '(' || arr[j] == '[' || arr[j] == '{' {
					stack = append(stack, arr[j])
				} else {
					switch arr[j] {
					case ')':
						if stack[len(stack)-1] == '(' {
							stack = stack[0 : len(stack)-1]

						}
					case ']':
						if stack[len(stack)-1] == '[' {
							stack = stack[0 : len(stack)-1]

						}
					case '}':
						if stack[len(stack)-1] == '{' {
							stack = stack[0 : len(stack)-1]
						}
					}
				}
			}

			if len(stack) == 0 {

				fmt.Println("OK")
			} else {

				fmt.Println("Error")
			}

		}
	}
}
