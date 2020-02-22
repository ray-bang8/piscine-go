package main

func ListReverse(l *List) {
	var CurrentPrev *NodeL
	var CurrentNext *NodeL
	CurrentNode := l.Head
	for CurrentNode != nil {
		CurrentNext = CurrentNode.Next
		CurrentNode.Next = CurrentPrev
		CurrentPrev = CurrentNode
		CurrentNode = CurrentNext

	}
	l.Head = CurrentPrev
	l.Tail = nil
}
func pushBack(l *List, x interface{}) {
	n := &NodeL{Data: x}

	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

// func main() {
// 	num1 := &piscine.NodeAddL{Num: 1}
// 	num1 = pushBack(num1, 3)
// 	num1 = pushBack(num1, 2)
// 	num1 = pushBack(num1, 4)
// 	num1 = pushBack(num1, 5)

// 	result := piscine.Reverse(num1)
// 	for tmp := result; tmp != nil; tmp = tmp.Next {
// 		fmt.Print(tmp.Num)
// 		if tmp.Next != nil {
// 			fmt.Print(" -> ")
// 		}
// 	}
// 	fmt.Println()
// }
