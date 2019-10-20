package main

import "fmt"

func BasicAtoi2(s string) int {
	res := 0
	for _, val := range s {
		a := 0
		if val >= '0' && val <= '9' {
			for i := '1'; i <= val; i++ {
				a++
			}
			res = res*10 + a
		} else {
			res = 0
			return res
		}
	}
	return res
}

func main() {
		s := "12345"
		s2 := "0000000012345"
		s3 := "012 345"
		s4 := "Hello World!"
		s5 := "+1234"
		s6 := "-1234"
		s7 := "++1234"
		s8 := "--1234"

		n := BasicAtoi2(s)
		n2 := BasicAtoi2(s2)
		n3 := BasicAtoi2(s3)
		n4:= BasicAtoi2(s4)
		n5 := BasicAtoi2(s5)
		n6 := BasicAtoi2(s6)
		n7 := BasicAtoi2(s7)
		n8 := BasicAtoi2(s8)

		fmt.Println(n)
		fmt.Println(n2)
		fmt.Println(n3)
		fmt.Println(n4)
		fmt.Println(n5)
		fmt.Println(n6)
		fmt.Println(n7)
		fmt.Println(n8)

	}
