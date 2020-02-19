package main

import (
	"fmt"
	"os"
)

func main() {

	arg := os.Args[1:]

	if len(arg) == 1 {
		str := arg[0]
		var count int
		var arr []int
		check := true

		for len(str) > 0 {
			count, str = counter(rune(str[0]), str)
			arr = append(arr, count)
		}
		for i, v := range arr {
			for j, k := range arr {
				if v == k && i != j {

					check = false
					break
				}
			}

		}
		fmt.Println(check)
	} else {
		fmt.Println()
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
func myAppend(arr []string, newString string) []string {
	count := 0
	for range arr {
		count++
	}

	newArr := make([]string, count+1)
	for index, element := range arr {
		newArr[index] = element
	}

	newArr[count] = newString
	return newArr
}
