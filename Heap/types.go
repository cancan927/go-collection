package Heap

type Ordered interface {
	Less(other Ordered) bool
}
