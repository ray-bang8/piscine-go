package main

import (
	"fmt"
	"os"
	"strconv"
)

func pop(stack []int) (int, []int) {
	// fmt.Println(stack[len(stack)-1], stack[:len(stack)-1])
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func RpnCalc(s []string) int {
	// ops -> + - * / %
	stackD := []int{}
	first := 0
	last := 0

	for _, v := range s {
		for _, vv := range v {
			switch vv {

			case '+':
				last, stackD = pop(stackD)
				first, stackD = pop(stackD)
				stackD = append(stackD, first+last)

			case '-':
				last, stackD = pop(stackD)
				first, stackD = pop(stackD)
				stackD = append(stackD, first-last)

			case '*':
				last, stackD = pop(stackD)
				first, stackD = pop(stackD)
				stackD = append(stackD, first*last)

			case '/':
				last, stackD = pop(stackD)
				first, stackD = pop(stackD)
				stackD = append(stackD, first/last)

			case '%':
				last, stackD = pop(stackD)
				first, stackD = pop(stackD)
				stackD = append(stackD, first%last)

			default:
				num, err := strconv.Atoi(string(vv))
				if err != nil {
					fmt.Println("Error!")
				} else {
					stackD = append(stackD, num)
				}
			}

		}

	}
	result := stackD[0]

	return result

}

func main() {
	arg := os.Args[1:]
	res := RpnCalc(arg)
	fmt.Println(res)
}
