package datastruct

import "container/heap"

//Len returns the length of the heap
func (h *Heap) Len() int {
	return h.length
}

//Less returns true if the element at index i is less than the element at index j
func (h *Heap) Less(i, j int) bool {
	ii, _ := h.Index(h.length - i - 1)
	ji, _ := h.Index(h.length - j - 1)
	in := ii.Value.(*HeapNode)
	jn := ji.Value.(*HeapNode)
	//use priority to compare and ismax to decide whether it is a max heap or min heap
	return (in.Priority > jn.Priority) == h.ismax
}

//Swap swaps the elements at index i and j
func (h *Heap) Swap(i, j int) {
	ii, _ := h.Index(h.length - i - 1)
	ji, _ := h.Index(h.length - j - 1)
	//we only swap the value, not the whole node
	ii.Value, ji.Value = ji.Value, ii.Value
}

//Push pushes the element x onto the heap.
func (h *Heap) Push(x interface{}) {
	node := NewListNode(x)
	h.Insert(node)
}

//Pop removes the element with the highest priority from the heap and returns it.
func (h *Heap) Pop() interface{} {
	if h.length == 0 {
		return nil
	}
	node, _ := h.Index(0)
	h.DeleteByIndex(0)
	return node.Value
}

//NewHeap returns a new Heap
func NewHeap(ismax bool) *Heap {
	return &Heap{
		LinkedList: NewLinkedList(),
		ismax:      ismax,
	}
}

//NewHeapNode returns a new HeapNode
func NewHeapNode(value interface{}, priority int) *HeapNode {
	return &HeapNode{
		Value:    value,
		Priority: priority,
	}
}

//SetIsMax sets the ismax field of the Heap, it will make a max heap become a min heap, when ismax is false, it will be a min heap and when ismax is true, it will be a max heap. We will reinit the heap after setting it.
func (h *Heap) SetIsMax(ismax bool) {
	h.ismax = ismax
	//reinit the heap
	heap.Init(h)
}
