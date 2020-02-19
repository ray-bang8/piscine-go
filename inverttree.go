import "fmt"

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

func InvertTree(root *TNode) *TNode {
	// проверка есть ли потомки
	if root == nil {
		return nil
	}

	// заносим во временную переменную левую ветку
	var t *TNode
	t = root.Left
	// рекурсивно вызываем для каждой подветки, т.е. слева(правое) справа(левое)
	root.Left = InvertTree(root.Right)
	fmt.Println("root.Left", root.Left)
	root.Right = InvertTree(t)
	return root
}

func BTreeInsertData(root *TNode, val int) *TNode {

	if root == nil {
		return &TNode{Val: val}
	}

	if val < root.Val {
		root.Left = BTreeInsertData(root.Left, val)
	} else {
		root.Right = BTreeInsertData(root.Right, val)
	}
	return root
}

func BTreeApplyInOrder(root *TNode, f func(...interface{}) (int, error)) {

	if root == nil {
		return
	}
	BTreeApplyInOrder(root.Left, f)
	f(root.Val)
	BTreeApplyInOrder(root.Right, f)
}