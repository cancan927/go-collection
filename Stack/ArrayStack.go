package Stack

import "errors"

type ArrayStack[T any] struct {
	elements []T
	size     int
}

func NewArrayStack[T any](cap int) *ArrayStack[T] {
	if cap < 10 {
		cap = 10
	}
	stack := new(ArrayStack[T])
	stack.elements = make([]T, 0, 10)
	stack.size = 0
	return stack

}

func (a *ArrayStack[T]) Size() int {
	return a.size
}

func (a ArrayStack[T]) IsEmpty() bool {
	return a.size == 0
}

func (a ArrayStack[T]) Peek() (T, error) {
	if a.IsEmpty() {
		var t T
		return t, errors.New("stack is empty")
	}
	return a.elements[a.size-1], nil
}

// 压栈
func (a ArrayStack[T]) Push(t T) {
	a.elements = append(a.elements, t)
	a.size++
}

// 弹出栈顶元素
func (a ArrayStack[T]) Pop() (T, error) {
	if a.IsEmpty() {
		var t T
		return t, errors.New("stack is empty")
	}
	top := a.elements[a.size-1]
	a.elements = a.elements[:a.size-1]
	a.size--
	return top, nil
}
