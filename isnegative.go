package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {

	if nb >= 0 {
		fmt.PrintRune('F')
	} else {
		fmt.PrintRune('T')
	}
	fmt.PrintRune('\n')
}
