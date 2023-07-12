package List

type List[T any] interface {
	// Size return the length of list
	Size() int

	//IsEmpty return whether the list is empty
	IsEmpty() bool

	//Clear the method which is clear the list
	Clear()

	//Get return the element based on index
	//if it doesn't exist,return nil and error
	//if it does exist,return the element and nil
	Get(index int) (T, error)

	//Contains return the val does exist in this list or not
	Contains(val T) bool

	//Set the element in this list based on index
	//if failed,it will be return an error;
	//if successful, return nil
	Set(index int, newVal T) error

	//Insert an element to this list based on index
	//if failed,return an error
	//if successful,return nil
	Insert(index int, newVal T) error

	//Append a newVal to end of this list
	Append(newVal T)

	//Delete an element in this list based on index
	Delete(index int) error
}
