package Queue

import "errors"

type Node[T any] struct {
	next, prev *Node[T]
	value      T
}
type LinkedQueue[T any] struct {
	root Node[T]
	size int
}

func (l *LinkedQueue[T]) init() *LinkedQueue[T] {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.size = 0
	return l
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return new(LinkedQueue[T]).init()
}

func (l *LinkedQueue[T]) Size() int {
	return l.size
}

func (l *LinkedQueue[T]) IsEmpty() bool {
	return l.size == 0
}

// EnQueue 新元素加入队列尾部
func (l *LinkedQueue[T]) EnQueue(newVal T) {
	e := &Node[T]{value: newVal}
	//新元素的后继指向root
	e.next = &l.root
	//新元素前驱指向原队尾节点
	e.prev = l.root.prev
	//原队尾节点后继指向新元素
	l.root.prev.next = e
	//root的前驱指向新元素
	l.root.prev = e
	l.size++

}

func (l *LinkedQueue[T]) DeQueue() (T, error) {
	if l.size == 0 {
		var t T
		return t, errors.New("queue is empty")
	}
	//
	front := l.root.next
	//将第二个元素的前驱指向front的前驱
	front.next.prev = &l.root
	//将根节点的后继指向第二个元素
	l.root.next = front.next
	l.size--
	return front.value, nil
}

func (l *LinkedQueue[T]) Front() (T, error) {
	if l.size == 0 {
		var t T
		return t, errors.New("queue is empty")
	} else {
		return l.root.next.value, nil
	}
}

func (l *LinkedQueue[T]) Clear() {
	l.init()
}
