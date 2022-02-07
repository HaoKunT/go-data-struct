package datastruct

type LinkedList struct {
	Head   *ListNode
	length int
}

type ListNode struct {
	Next  *ListNode
	Value interface{}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Value interface{}
}

type BinaryTreeHandleFunc func(node *BinaryTreeNode) error
type BinaryTreeErrorHandleFunc func(err error)

// Stack is a linked list
type Stack struct {
	*LinkedList
}

type Queue struct {
	*LinkedList
}

//Heap is a list which implements heap interface, we use ismax to indicate whether the heap is a max heap or min heap
type Heap struct {
	*LinkedList
	ismax bool //true for max heap, false for min heap
}

//HeapNode is a node in Heap
type HeapNode struct {
	Value    interface{}
	Priority int
}
