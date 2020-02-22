package main

import (
	"fmt"
)

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {

	res := ListToInt(num1) + ListToInt(num2)
	str := Itoa(res)
	firstDigit := int(str[len(str)-1] - 48)
	list := &NodeAddL{Num: firstDigit}
	for i := len(str) - 2; i >= 0; i-- {
		list = pushFront(list, int(str[i]-48))
	}
	return list
}

func ListToInt(num *NodeAddL) int {
	str := ""
	current := num
	for current != nil {
		str += Itoa(current.Num)
		current = current.Next
	}
	nbr := Atoi(str)
	return nbr
}

func pushFront(node *NodeAddL, num int) *NodeAddL {
	a := &NodeAddL{
		Num:  num,
		Next: node,
	}
	return a
}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Itoa(n int) string {

	res := ""
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	return res
}
func Atoi(s string) int {
	min := 0
	plu := 0
	sum := 0
	for i, v := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for j := '0'; j < v; j++ {
				num++
			}
			sum = sum*10 + num
		} else if s[i] == '-' && i == 0 {
			min++
		} else if s[i] == '+' && i == 0 {
			plu++
		} else {
			return 0
		}
	}
	if min == 1 {
		sum = -sum
	}
	if min != 1 && sum < 0 {
		return 0
	}
	if min == 1 && sum > 0 {
		return 0
	}
	return sum

}
func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num: 6}
	num1 = pushFront(num1, 2)
	num1 = pushFront(num1, 4)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num: 2}
	num2 = pushFront(num2, 1)
	num2 = pushFront(num2, 3)

	// 9 -> 0 -> 7
	result := AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
