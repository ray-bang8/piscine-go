package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := os.Args[1:]

	if len(arg) == 1 {
		str := arg[0]
		var count int
		var arr []int
		check := "true"

		for len(str) > 0 {
			count, str = counter(rune(str[0]), str)
			arr = Append(arr, count)
		}
		for i, v := range arr {
			for j, k := range arr {
				if v == k && i != j {

					check = "false"
					break
				}
			}

		}
		for i := range check {
			z01.PrintRune(rune(check[i]))
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}

}

func counter(c rune, s string) (int, string) {
	count := 0
	str := ""
	for _, v := range s {
		if v == c {
			count++

		} else {
			str = str + string(v)
		}
	}
	return count, str
}

func Append(narr []int, n int) []int {

	newarr := make([]int, len(narr)+1)

	for i, lol := range narr {
		newarr[i] = lol
	}

	newarr[len(narr)] = n
	return newarr

}
