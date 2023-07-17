package Queue

import "errors"

type ArrayQueue[T any] struct {
	elements []T
	len      int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	if capacity < 10 {
		capacity = 10
	}
	queue := new(ArrayQueue[T])
	queue.elements = make([]T, 0, capacity)
	queue.len = 0
	return queue
}

func (a *ArrayQueue[T]) Size() int {
	return a.len
}

func (a *ArrayQueue[T]) IsEmpty() bool {
	return a.len == 0
}

func (a *ArrayQueue[T]) EnQueue(newVal T) {
	a.elements = append(a.elements, newVal)
	a.len++
}

func (a *ArrayQueue[T]) DeQueue() (T, error) {
	if a.Size() == 0 {
		var t T
		return t, errors.New("queue is empty")
	} else {
		result := a.elements[0]
		a.elements = a.elements[1:]
		a.len--
		return result, nil
	}
}

func (a *ArrayQueue[T]) Front() (T, error) {
	if a.len == 0 {
		var t T
		return t, errors.New("queue is empty")
	}
	return a.elements[0], nil
}

func (a *ArrayQueue[T]) Clear() {
	a.elements = make([]T, 0, 10)
	a.len = 0
}
