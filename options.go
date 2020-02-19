package main

import (
	"fmt"
	"os"
)

func ErrorHandle(i int) {
	switch i {
	case 0:
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
	case 1:
		fmt.Println("Invalid Option")
	}
}

func Replace(r rune, pattern string) string {
	index := 0
	res := ""
	sample := "******zy xwvutsrq ponmlkji hgfedcba"
	for i, v := range sample {
		if r == v {
			index = i
		}
	}
	for j, k := range pattern {
		if j == index {
			k = '1'
		}
		res = res + string(k)
	}
	return res
}

func main() {
	arg := os.Args[1:]
	pattern := "00000000 00000000 00000000 00000000"
	var flag bool
	if len(arg) < 1 {
		ErrorHandle(0)
	}
	for _, vv := range arg {
		for i, v := range vv {
			if i == 0 && v == '-' {
				flag = true
			}
			if vv[i] == '-' && vv[i+1] == 'h' && i < len(vv) {
				ErrorHandle(0)
				return
			}
			if flag {
				if v >= 'a' && v <= 'z' {
					res := Replace(v, pattern)
					pattern = res
				} else if v == '-' {
				} else {
					ErrorHandle(1)
					return
				}
			}
		}
	}
	fmt.Println(pattern)
}
