package Tree

/*
*
经典的二叉树，所有元素应该是严格不相等的，且元素之间必须支持<,>的比较操作
*/

// GenericTree 结构体存储树的根节点和树的大小
type GenericTree[T Ordered] struct {
	root gtreeNode[T]
	size int
}

// gtreeNode 结构体用于存储树的节点
type gtreeNode[T Ordered] struct {
	left, right *gtreeNode[T]
	value       T
}

// insertNode 方法将一个元素插入到树中
func (b *GenericTree[T]) insertNode(node *gtreeNode[T], val T) {
	if node == nil {
		// 如果节点为空，将新节点插入到这个位置
		*node = gtreeNode[T]{value: val}
		b.size++
		return
	}
	// 如果节点不为空，根据值的大小将新节点插入到左子树或右子树中
	if val.Less(node.value) {
		// 如果值小于节点的值，将新节点插入到左子树中
		b.insertNode(node.left, val)
		// 如果值大于节点的值，将新节点插入到右子树中
	} else if val.Greater(node.value) {
		b.insertNode(node.right, val)
	}
}

// containsNode 方法返回一个布尔值，指示树中是否包含给定的元素
func (b *GenericTree[T]) containsNode(node *gtreeNode[T], val T) bool {
	if node == nil {
		return false
	}
	if node.value.Equal(val) {
		return true
	}
	if val.Less(node.value) {
		return b.containsNode(node.left, val)
	}
	return b.containsNode(node.right, val)
}

// NewBinarySearchTree 方法创建一个新的二叉树
func NewBinarySearchTree[T Ordered]() *GenericTree[T] {
	tree := &GenericTree[T]{size: 0}
	return tree
}

// Size 方法返回树的大小
func (b *GenericTree[T]) Size() int {
	return b.size
}

// IsEmpty 方法返回一个布尔值，指示树是否为空
func (b *GenericTree[T]) IsEmpty() bool {
	return b.size == 0
}

// Clear 方法清除树的元素并将大小设置为0
func (b *GenericTree[T]) Clear() {
	b = NewBinarySearchTree[T]()
}

// Insert 方法将一个元素插入到树中
func (b *GenericTree[T]) Insert(val T) {
	if b.size == 0 {
		b.root = gtreeNode[T]{value: val}
		b.size++
		return
	}
	b.insertNode(&b.root, val)
}

// Contains 方法返回一个布尔值，指示树中是否包含给定的元素
func (b *GenericTree[T]) Contains(val T) bool {
	return b.containsNode(&b.root, val)
}

// Delete 方法从树中删除一个元素
func (b *GenericTree[T]) Delete(val T) {
	b.deleteNode(&b.root, val)
}

// deleteNode 方法从树中删除一个元素
func (b *GenericTree[T]) deleteNode(node *gtreeNode[T], val T) *gtreeNode[T] {
	if node == nil {
		return node
	}
	if node.value.Greater(val) {
		node.left = b.deleteNode(node.left, val)
	} else if node.value.Less(val) {
		node.right = b.deleteNode(node.right, val)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		node.value = b.minValue(node.right)
		node.right = b.deleteNode(node.right, node.value)
	}
	return node
}

// minValue 方法返回树中的最小值
func (b *GenericTree[T]) minValue(node *gtreeNode[T]) T {
	minValue := node.value
	for node.left != nil {
		minValue = node.left.value
		node = node.left
	}
	return minValue
}
