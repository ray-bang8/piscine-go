package main

import (
	"fmt"
	"os"
	"strconv"
)

func ISPOWEROFF() {
	if len(os.Args) == 2 {
		x, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for x%2 == 0 {
				if x == 0 {
					break
				}
				x = x / 2
			}
			if x == 1 {
				fmt.Println(true)
			}
			if x == 0 || x != 1 {
				fmt.Println(false)
			}
		}
	} else {
		fmt.Println()
	}
}
func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		x, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println(err.Error())
		} else if x != 0 && (x&(x-1)) == 0 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}

	} else {
		fmt.Println()
	}
}
