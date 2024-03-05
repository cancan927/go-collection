package Tree

// 构建一个搜索二叉树
type BSTreeNode struct {
	Value int
	Left  *BSTreeNode
	Right *BSTreeNode
}

func (n *BSTreeNode) Insert(value int) {
	if n == nil {
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &BSTreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &BSTreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *BSTreeNode) Search(value int) *BSTreeNode {
	if n == nil || n.Value == value {
		return n
	}

	if value < n.Value {
		return n.Left.Search(value)
	}

	return n.Right.Search(value)
}
