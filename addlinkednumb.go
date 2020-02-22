

import "strconv"

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {

	res := ListToInt(num1) + ListToInt(num2)
	str := strconv.Itoa(res)
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
		str += strconv.Itoa(current.Num)
		current = current.Next
	}
	nbr, _ := strconv.Atoi(str)
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