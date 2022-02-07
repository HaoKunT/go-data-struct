package datastruct

import (
	"bytes"
	"errors"
	"fmt"
)

func NewLinkedList() *LinkedList {
	ll := new(LinkedList)
	return ll
}

func (ll *LinkedList) IndexOf(node *ListNode) (int, error) {
	var index int
	var tmpNode = ll.Head
	for {
		if index >= ll.length {
			return -1, errors.New("No Node")
		}
		if node == tmpNode {
			return index, nil
		} else {
			tmpNode = tmpNode.Next
			index++
		}
	}
}

func (ll *LinkedList) Index(index int) (*ListNode, error) {
	if index < 0 || index >= ll.length {
		return nil, fmt.Errorf("out of index: %d is out of bounds, it should be between [0, %d)", index, ll.length)
	}
	if index == 0 {
		return ll.Head, nil
	}
	tmpNode := ll.Head
	for ; index > 0; index-- {
		tmpNode = tmpNode.Next
	}
	return tmpNode, nil
}

func (ll *LinkedList) Insert(node *ListNode) {
	node.Next = ll.Head
	ll.Head = node
	ll.length++
}

func (ll *LinkedList) InsertAt(node *ListNode, index int) error {
	if index < 0 || index > ll.length {
		return fmt.Errorf("index out of range: [0, %d], get %d", ll.length, index)
	}
	tmpNode := ll.Head
	// if index == 0, it equals Insert
	if index == 0 {
		ll.Insert(node)
		return nil
	}
	index--
	for ; index > 0; index-- {
		tmpNode = tmpNode.Next
	}
	// now tmpNode point to ll[index-1]
	node.Next = tmpNode.Next
	tmpNode.Next = node
	ll.length++
	return nil
}

//Pop will return the first node and delete it from the list
func (ll *LinkedList) Pop() (*ListNode, error) {
	if ll.length == 0 {
		return nil, errors.New("No Node")
	}
	tmpNode := ll.Head
	ll.Head = ll.Head.Next
	tmpNode.Next = nil
	ll.length--
	return tmpNode, nil
}

func (ll *LinkedList) String() string {
	buf := bytes.NewBuffer([]byte{})
	tmpNode := ll.Head
	for i := 0; i < ll.length; i++ {
		buf.WriteString(fmt.Sprintf("%v", tmpNode.Value))
		if i != ll.length-1 {
			buf.WriteString("->")
		}
		tmpNode = tmpNode.Next
	}
	return buf.String()
}

func (ll *LinkedList) DeleteByIndex(index int) error {
	if index < 0 || index >= ll.length {
		return fmt.Errorf("index out of range: [0, %d), get %d", ll.length, index)
	}
	if index == 0 {
		ll.Head = ll.Head.Next
		ll.length--
		return nil
	}
	lastNode, _ := ll.Index(index - 1)
	deleteNode := lastNode.Next
	lastNode.Next = deleteNode.Next
	deleteNode.Next = nil
	ll.length--
	return nil
}

func (ll *LinkedList) DeleteByNode(node *ListNode) error {
	index, err := ll.IndexOf(node)
	if err != nil {
		return err
	}
	err = ll.DeleteByIndex(index)
	if err != nil {
		return err
	}
	return nil
}

func (ll *LinkedList) Modify(index int, value interface{}) error {
	node, err := ll.Index(index)
	if err != nil {
		return err
	}
	node.Value = value
	return nil
}

func (ll *LinkedList) GetFirst(value interface{}) (int, *ListNode) {
	tmpNode := ll.Head
	for i := 0; i < ll.length; i++ {
		if tmpNode.Value == value {
			return i, tmpNode
		}
		tmpNode = tmpNode.Next
	}
	return -1, nil
}

func (ll *LinkedList) GetLength() int {
	return ll.length
}

func (ll *LinkedList) Clear() {
	// use simple method
	ll.Head = nil
	ll.length = 0
}

func NewListNode(value interface{}) *ListNode {
	ln := new(ListNode)
	ln.Value = value
	return ln
}
