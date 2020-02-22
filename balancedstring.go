

import (
	"fmt"
	"os"
)

func BALANCESTRING() {
	input := os.Args[1:]
	if len(input) != 1 {
		fmt.Println()
		return
	}
	count := 0
	bims := 0
	for _, val := range input[0] {
		if val == 'C' {
			count++
		} else {
			count--
		}
		if count == 0 {
			bims++
		}
	}
	fmt.Println(bims)
}
