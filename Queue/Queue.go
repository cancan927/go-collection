package Queue

type Queue[T any] interface {
	// Size 队列元素个数
	Size() int

	// IsEmpty 判断队列是否为空
	IsEmpty() bool

	// EnQueue 入队
	EnQueue(T)

	// DeQueue 出队
	DeQueue() (T, error)

	// Front 返回队头元素
	Front() (T, error)

	// Clear 清空队列
	Clear()
}
