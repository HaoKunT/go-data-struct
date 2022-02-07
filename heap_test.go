package datastruct

import (
	"container/heap"
	"testing"
)

//TestNewMaxHeapNode
func TestNewHeapNode(t *testing.T) {
	node1 := NewHeapNode(1, 1)

	if node1.Value != 1 || node1.Priority != 1 {
		t.Errorf("HeapNode: NewHeapNode error, it should be \" node1.Value == 1 && node1.Priority == 1\"")
	}
}

//TestNewMaxHeap
func TestNewMaxHeap(t *testing.T) {
	heap := NewHeap(true)

	if heap.length != 0 {
		t.Errorf("Heap: NewMaxHeap error, it should be \" heap.length == 0\"")
	}
}

//TestHeap
func TestHeap(t *testing.T) {
	h := NewHeap(true)
	heap.Init(h)

	heap.Push(h, &HeapNode{Value: 3, Priority: 3})
	heap.Push(h, &HeapNode{Value: 2, Priority: 2})
	heap.Push(h, &HeapNode{Value: 1, Priority: 1})
	heap.Push(h, &HeapNode{Value: 4, Priority: 4})
	heap.Push(h, &HeapNode{Value: 5, Priority: 5})
	heap.Push(h, &HeapNode{Value: 6, Priority: 6})

	if h.length != 6 {
		t.Errorf("Heap: MaxHeap error, it should be \" h.length == 6\" but got %d", h.length)
	}

	v := heap.Pop(h).(*HeapNode).Value
	if v != 6 {
		t.Errorf("Heap: MaxHeap error, it should be \"  h.Pop().(*HeapNode).Value == 6\" but got %d", v)
	}

	h.SetIsMax(false)
	v = heap.Pop(h).(*HeapNode).Value
	if v != 1 {
		t.Errorf("Heap: MinHeap error, it should be \"  h.Pop().(*HeapNode).Value == 1\" but got %d", v)
	}

}

//TestHeapReversed
func TestHeapReversed(t *testing.T) {
	h := NewHeap(true)
	heap.Init(h)

	heap.Push(h, &HeapNode{Value: 1, Priority: 4})
	heap.Push(h, &HeapNode{Value: 2, Priority: 3})
	heap.Push(h, &HeapNode{Value: 3, Priority: 2})
	heap.Push(h, &HeapNode{Value: 4, Priority: 1})

	if h.length != 4 {
		t.Errorf("Heap: MaxHeap error, it should be \" h.length == 5\" but got %d", h.length)
	}

	v := heap.Pop(h).(*HeapNode).Value
	if v != 1 {
		t.Errorf("Heap: MaxHeap error, it should be \"  h.Pop().(*HeapNode).Value == 1\" but got %d", v)
	}

	h.SetIsMax(false)
	v = heap.Pop(h).(*HeapNode).Value
	if v != 4 {
		t.Errorf("Heap: MinHeap error, it should be \"  h.Pop().(*HeapNode).Value == 4\" but got %d", v)
	}
}
