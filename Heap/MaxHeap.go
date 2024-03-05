package Heap

type MaxHeap struct {
	heap     []int
	heapSize int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap:     make([]int, 0),
		heapSize: 0,
	}
}

func (h *MaxHeap) Size() int {
	return h.heapSize
}
func (h *MaxHeap) IsEmpty() bool {
	return h.heapSize == 0
}

// 新增元素
func (h *MaxHeap) Push(v int) {
	if h.heapSize == len(h.heap) {
		h.heap = append(h.heap, v)
	} else {
		h.heap[h.heapSize] = v
	}
	h.heapInsert(h.heapSize)
	h.heapSize++
}

func (h MaxHeap) Pop() int {
	ans := h.heap[0]
	h.heapSize--
	if h.heapSize > 0 {
		h.heap[0], h.heap[h.heapSize] = h.heap[h.heapSize], h.heap[0]
		h.heapIfy(0, h.heapSize)
	}
	return ans
}

// 新来的数，在i位置
// 不停往上，如果大于父亲，交换
// 到了根节点（0位置），或者不再大于父亲，停
func (h *MaxHeap) heapInsert(i int) {
	for h.heap[i] > h.heap[(i-1)/2] {
		h.heap[i], h.heap[(i-1)/2] = h.heap[(i-1)/2], h.heap[i]
		i = (i - 1) / 2
	}
}

func (h *MaxHeap) heapIfy(i int, size int) {

}
