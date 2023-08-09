package Tree

type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 先序遍历
func preOrder(head *TreeNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	result = append(result, head.value)
	preOrder(head.left)
	preOrder(head.right)
	return result
}

// 中序遍历
func inOrder(head *TreeNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	inOrder(head.left)
	result = append(result, head.value)
	inOrder(head.right)
	return result
}

func postOrder(head *TreeNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	postOrder(head.left)
	postOrder(head.right)
	result = append(result, head.value)
	return result
}
