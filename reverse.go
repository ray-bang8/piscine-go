package main

import "fmt"

// func ListReverse(l *List) {
// 	var CurrentPrev *NodeL
// 	var CurrentNext *NodeL
// 	CurrentNode := l.Head
// 	for CurrentNode != nil {
// 		CurrentNext = CurrentNode.Next
// 		CurrentNode.Next = CurrentPrev
// 		CurrentPrev = CurrentNode
// 		CurrentNode = CurrentNext

// 	}
// 	l.Head = CurrentPrev
// 	l.Tail = nil
// }
func pushBack(n *NodeAddL, x int) {
	n := &NodeAddL{Num: x}

}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)
	fmt.Println(num1)
}
