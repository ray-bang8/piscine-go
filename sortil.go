func ListSort(l *NodeAddL) *NodeAddL {

	for i := l; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}
	return l
}
______________
func Sortll(node *NodeAddL) *NodeAddL {
	copy := node

	for x := copy; x.Next != nil; x = x.Next {
		for it := copy; it.Next != nil; it = it.Next {
			if it.Num < it.Next.Num {
				it.Num, it.Next.Num = it.Next.Num, it.Num
			}
		}
	}

	return copy
}