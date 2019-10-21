package main

import (
	"fmt"
)

func RecursivePower(nb int, power int) int {
	if power < 0 || power > 20 {
		return 1
	}
	if power == 0 {
		return 1
	} else if power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return nb
}

func main() {
	arg1 := 5
	arg2 := 2
	fmt.Println(RecursivePower(arg1, arg2))
}
