package Stack

type Stack[T any] interface {
	Size() int

	IsEmpty() bool

	Peek() (T, error)

	Push(T)

	Pop() (T, error)
}
