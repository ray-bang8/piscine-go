package main

import "fmt"

func BasicAtoi(s string) int {
	res := 0
	for _, val := range s {
		a := 0
		for i := '1'; i <= val; i++ {
			a++
		}
		res = res*10 + a
	}
	return res
}
func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "000000"

	n := BasicAtoi(s)
	n2 := BasicAtoi(s2)
	n3 := BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)

}