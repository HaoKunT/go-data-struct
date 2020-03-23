package dataStruct

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
