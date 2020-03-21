package dataStruct

type LinkedList struct {
	Head   *ListNode
	length int
}

type ListNode struct {
	Next  *ListNode
	Value interface{}
}
