package Stack

import "errors"

type LinkedStack[T any] struct {
	root node[T]
	size int
}

type node[T any] struct {
	next, prev *node[T]
	value      T
}

func (l *LinkedStack[T]) init() *LinkedStack[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.size = 0
	return l
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return new(LinkedStack[T]).init()

}
func (l *LinkedStack[T]) Size() int {
	return l.size
}

func (l *LinkedStack[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedStack[T]) Peek() (T, error) {
	if l.size == 0 {
		var t T
		return t, errors.New("stack is empty")
	}
	//相当于获取链表的最后一个元素
	return l.root.prev.value, nil
}

func (l *LinkedStack[T]) Push(t T) {
	newNode := &node[T]{value: t}

	//先连后断
	newNode.prev = l.root.prev
	newNode.next = &l.root

	//
	l.root.prev.next = newNode
	l.root.prev = newNode

	l.size++
}

func (l LinkedStack[T]) Pop() (T, error) {
	if l.IsEmpty() {
		var t T
		return t, errors.New("stack is empty")
	}
	top := l.root.prev
	value := top.value

	//先连
	top.prev.next = top.next
	top.next.prev = top.prev

	top.prev = nil
	top.next = nil

	l.size--
	return value, nil
}
