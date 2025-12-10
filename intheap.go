package heaputill

import "container/heap"

// IntHeap is a heap of ints that can behave as either a min-heap or max-heap
// depending on the comparator used.
type IntHeap struct {
	data []int
	less func(a, b int) bool // a "comes before" b if less(a, b) == true
}

// Ensure *IntHeap implements heap.Interface.
var _ heap.Interface = (*IntHeap)(nil)

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}

// NewMinIntHeap constructs a min-heap with optional initial values.
func NewMinIntHeap(nums ...int) *IntHeap {
	h := &IntHeap{
		less: func(a, b int) bool { return a < b }, // min-heap
	}
	h.data = append(h.data, nums...)
	heap.Init(h)
	return h
}

// NewMaxIntHeap constructs a max-heap with optional initial values.
func NewMaxIntHeap(nums ...int) *IntHeap {
	h := &IntHeap{
		less: func(a, b int) bool { return a > b }, // max-heap
	}
	h.data = append(h.data, nums...)
	heap.Init(h)
	return h
}

func (h *IntHeap) PushInt(x int) {
	heap.Push(h, x)
}

func (h *IntHeap) PopInt() int {
	return heap.Pop(h).(int)
}

// Peek returns the root element (min for min-heap, max for max-heap).
// Caller must ensure h.Len() > 0.
func (h *IntHeap) Peek() int {
	return h.data[0]
}

// Data exposes the underlying slice (read-only use recommended).
func (h *IntHeap) Data() []int {
	return h.data
}
