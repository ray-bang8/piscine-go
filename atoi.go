package main

import "fmt"

func Atoi(s string) int {
	nb := S(s)
	meow := []rune(s)
	s = ""
	final := 0
	plus := 0
	minus := 0
	num := 0
	for i := 0; i < nb; i++ {
		if meow[i] >= '0' && meow[i] <= '9' {
			num = 0
			for j := '0'; j < meow[i]; j++ {
				num++
			}
			final = final*10 + num
		} else if meow[i] == '+' && i == 0 {
			plus += 1
		} else if meow[i] == '-' && i == 0 {
			minus += 1

		} else {
			return 0
		}
	}
	if minus == 1 {
		final = -final
	}
	return final
}
func S(s string) int {
	a := 0
	for range s {
		a++
	}
	return a
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

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
