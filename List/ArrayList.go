package List

import (
	"errors"
	"reflect"
)

// ArrayList struct is used to store an array of elements and the size of the array
type ArrayList[T any] struct {
	// This is the elements of the array
	elements []T
	// This is the size of the array
	size int
}

// NewArrayList create a new ArrayList with a capacity
func NewArrayList[T any](capacity int) *ArrayList[T] {
	// Check if the capacity is less than 10
	if capacity < 10 {
		// If so, set the capacity to 10
		capacity = 10
	}
	// Create a new ArrayList
	list := &ArrayList[T]{
		// Initialize the elements array with an empty slice of type T
		elements: make([]T, 0, capacity),
		// Initialize the size to 0
		size: 0,
	}
	// Return the new ArrayList
	return list
}

// Size function returns the size of the ArrayList
func (a *ArrayList[T]) Size() int {
	return a.size
}

// IsEmpty function returns a boolean value indicating whether the ArrayList is empty or not
func (a *ArrayList[T]) IsEmpty() bool {
	return a.size == 0
}

// Clear function clears the elements of an ArrayList[T] and sets the size to 0
func (a *ArrayList[T]) Clear() {
	// Create a new slice of type T with the same length as the original array
	a.elements = make([]T, 0, 10)
	// Set the size of the ArrayList[T] to 0
	a.size = 0
}

// Get function takes in an ArrayList of type T and an integer index and returns the element at the given index and an error
func (a *ArrayList[T]) Get(index int) (T, error) {
	//Check if the index is within the bounds of the ArrayList
	if index < 0 || index >= a.size {
		//If not, return an error
		var t T //temp variable,represent zero-value of type T
		return t, errors.New("index out of bounds")
	}
	//If the index is within the bounds, return the element at the given index
	return a.elements[index], nil
}

// Contains function checks if the array list contains a given value
func (a *ArrayList[T]) Contains(val T) bool {
	// Iterate through the elements in the array list
	for _, element := range a.elements {
		// If the element matches the value, return true
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
