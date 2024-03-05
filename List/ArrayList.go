package List

import (
	"errors"
	"reflect"
)

// ArrayList 结构体用于存储数组列表的元素和大小
type ArrayList[T any] struct {
	elements []T
	size     int
}

// NewArrayList 方法创建一个新的数组列表
func NewArrayList[T any](capacity int) *ArrayList[T] {
	// 检查容量是否小于10
	if capacity < 10 {
		// 如果是，将容量设置为10
		capacity = 10
	}
	// 创建一个新的ArrayList[T]，并将其elements数组初始化为一个空的T类型切片
	list := &ArrayList[T]{
		// 初始化
		elements: make([]T, 0, capacity),
		size:     0,
	}
	// 返回新的ArrayList[T]
	return list
}

// Size 方法返回数组列表的大小
func (a *ArrayList[T]) Size() int {
	return a.size
}

// IsEmpty 方法返回一个布尔值，指示数组列表是否为空
func (a *ArrayList[T]) IsEmpty() bool {
	return a.size == 0
}

// Clear 方法清除数组列表的元素并将大小设置为0
func (a *ArrayList[T]) Clear() {
	// 创建一个新的T类型切片，并将其赋值给elements
	a.elements = make([]T, 0, 10)
	// 将size设置为0
	a.size = 0
}

// Get 方法返回数组列表中给定索引的元素
func (a *ArrayList[T]) Get(index int) (T, error) {
	//检查索引是否在数组列表的范围内
	if index < 0 || index >= a.size {
		//如果不在范围内，返回一个错误
		var t T //临时变量，代表T类型的零值
		return t, errors.New("index out of bounds")
	}
	//If the index is within the bounds, return the element at the given index
	return a.elements[index], nil
}

// Contains 方法返回一个布尔值，指示数组列表是否包含给定的值
func (a *ArrayList[T]) Contains(val T) bool {
	// 遍历数组列表的元素
	for _, element := range a.elements {
		// 如果元素与值匹配，则返回true
		if reflect.DeepEqual(val, element) {
			return true
		}
	}
	// If the element does not match the value, return false
	return false
}

// Set function sets a value in an array list at a given index
func (a *ArrayList[T]) Set(index int, newVal T) error {
	// Check if the index is within the bounds of the array list
	if index < 0 || index >= a.size {
		// If not, return an error
		return errors.New("index out of bounds")
	}
	// Set the value at the given index in the array list
	a.elements[index] = newVal
	// Return nil
	return nil
}

// Append a new value to the ArrayList
func (a *ArrayList[T]) Append(newVal T) {
	// Append the new value to the elements of the ArrayList
	a.elements = append(a.elements, newVal)
	// Increment the size of the ArrayList
	a.size++
}

// Insert function inserts a new value into an array list at a given index
func (a *ArrayList[T]) Insert(index int, newVal T) error {
	//Check if the index is within the bounds of the array list
	if index < 0 || index > a.size {
		return errors.New("index out of bounds")
	}
	// consider 2 spacial cases:
	//1. insert into the head
	if index == 0 {
		//Create a new array list with the size of the array list + 10
		newElements := make([]T, 0, a.size+10)
		//Append the new value to the new array list
		newElements = append(newElements, newVal)
		//Append the elements of the array list to the new array list
		a.elements = append(newElements, a.elements[:]...)
		//Increment the size of the array list
		a.size++
		//Return nil
		return nil
	}
	//2. insert into the tail
	if index == a.size {
		//Append the new value to the array list
		a.elements = append(a.elements, newVal)
		//Return nil
		return nil
	}
	//Normal case
	//Create a new array list with the size of the array list + 10
	newElements := make([]T, 0, a.size+10)
	//Append the elements of the array list before the given index to the new array list
	newElements = append(newElements, a.elements[:index]...)
	//Append the new value to the new array list
	newElements = append(newElements, newVal)
	//Append the elements of the array list after the given index to the new array list
	newElements = append(newElements, a.elements[index:]...)
	//Set the elements of the array list to the new array list
	a.elements = newElements
	//Increment the size of the array list
	a.size++
	//Return nil
	return nil
}

// Delete an element from the ArrayList at a given index
func (a *ArrayList[T]) Delete(index int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	a.size--
	return nil
}
