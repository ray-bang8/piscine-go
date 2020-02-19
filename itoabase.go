package main

import "fmt"

func ItoaBase(value, base int) string {
	res := ""
	for value > 0 {
		res = string(value%base+48) + res
		value /= base
	}
	return res

}

func main() {
	baseString := "0123456789ABCDEF"
	s := ItoaBase(123456789, len(baseString))
	fmt.Println(s)

}
