package List

import (
	"errors"
	"reflect"
)

type ArrayList[T any] struct {
	elements []T
	size     int
}

func NewArrayList[T any](capacity int) *ArrayList[T] {
	if capacity < 10 {
		capacity = 10
	}
	list := &ArrayList[T]{
		elements: make([]T, 0, capacity),
		size:     0,
	}
	return list
}

func (a *ArrayList[T]) Size() int {
	return a.size
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList[T]) Clear() {
	a.elements = make([]T, 0, 10)
	a.size = 0
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.size {
		var t T //temp variable,represent zero-value of type T
		return t, errors.New("index out of bounds")
	}
	return a.elements[index], nil
}

func (a *ArrayList[T]) Contains(val T) bool {
	for _, element := range a.elements {
		if reflect.DeepEqual(val, element) {
			return true
		}
	}
	return false
}

func (a *ArrayList[T]) Set(index int, newVal T) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.elements[index] = newVal
	return nil
}

func (a *ArrayList[T]) Insert(index int, newVal T) error {
	if index < 0 || index > a.size {
		return errors.New("index out of bounds")
	}
	// consider 2 spacial cases:
	// 1. insert into the head
	if index == 0 {
		newElements := make([]T, 0, a.size+10)
		newElements = append(newElements, newVal)
		a.elements = append(newElements, a.elements[:]...)
		a.size++
		return nil
	}
	//2. insert into the tail
	if index == a.size {
		a.Append(newVal)
		return nil
	}

	//normal case
	newElements := make([]T, 0, a.size+10)
	newElements = append(newElements, a.elements[:index]...)
	newElements = append(newElements, newVal)
	newElements = append(newElements, a.elements[index:]...)
	a.elements = newElements
	a.size++
	return nil
}

func (a *ArrayList[T]) Append(newVal T) {
	a.elements = append(a.elements, newVal)
	a.size++
}

func (a *ArrayList[T]) Delete(index int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	a.size--
	return nil
}
