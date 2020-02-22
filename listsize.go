package main

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
___________________
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		return
	} else {
		next := l.Head
		for next.Next != nil {
			next = next.Next
		}
		next.Next = n
	}
}
