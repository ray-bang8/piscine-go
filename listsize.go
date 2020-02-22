
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	t := l.Head
	c := 0
	for t != nil {
		c++
		t = t.Next
	}
	return c
}

func ListPushBack(l *List, x interface{}) {
	n := &NodeL{Data: x}

	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
