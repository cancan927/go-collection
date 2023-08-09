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
	left := preOrder(head.left)
	result = append(result, left...)
	right := preOrder(head.right)
	result = append(result, right...)
	return result
}

// 中序遍历
func inOrder(head *TreeNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	left := inOrder(head.left)
	result = append(result, left...)
	result = append(result, head.value)
	right := inOrder(head.right)
	result = append(result, right...)
	return result
}

func postOrder(head *TreeNode) []int {
	result := make([]int, 0)
	if head == nil {
		return result
	}
	left := postOrder(head.left)
	result = append(result, left...)
	right := postOrder(head.right)
	result = append(result, right...)
	result = append(result, head.value)
	return result
}
