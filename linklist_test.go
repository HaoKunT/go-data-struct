package datastruct

import (
	"testing"
)

func TestLinkedListIndexOf(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 2,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)
	if r, err := ll.IndexOf(node1); err != nil || r != 2 {
		t.Errorf("LinkedList: IndexOf error: %s, r should be 2 but got %d", err, r)
	}
	if r, err := ll.IndexOf(node2); err != nil || r != 1 {
		t.Errorf("LinkedList: IndexOf error: %s, r should be 1 but got %d", err, r)
	}
	if r, err := ll.IndexOf(node3); err != nil || r != 0 {
		t.Errorf("LinkedList: IndexOf error: %s, r should be 0 but got %d", err, r)
	}
	node4 := NewListNode(4)
	if r, err := ll.IndexOf(node4); err == nil || r != -1 {
		t.Errorf("LinkedList: IndexOf error: %s, r should be -1 but got %d", err, r)
	}
}

func TestLinkedListInsert(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)
}

func TestLinkedListInsertAt(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)
	var node4 = NewListNode(4)
	err := ll.InsertAt(node4, 2)
	if err != nil {
		t.Errorf("LinkedList: InsertAt error: %s", err)
	}
	if r := ll.String(); r != "3->2->4->1" {
		t.Errorf("LinkedList: String error: it should be 3->2->4->1 but get: %s", r)
	}
	if err := ll.InsertAt(NewListNode(5), 999); err == nil {
		t.Errorf("LinkedList: InsertAt error: it should be error, but get: %s", err)
	}
}

func TestLinkedListIndex(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	if r, err := ll.Index(1); r != node2 || err != nil {
		t.Errorf("LinkedList: Index error: %s, it should be %v, but get %v", err, node2, r)
	}
	if r, err := ll.Index(999); r != nil || err == nil {
		t.Errorf("LinkedList: Index error: %s, it should be %v, but get %v", err, nil, r)
	}
}

//TestLinkedListPop
func TestLinkedListPop(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	tNode, err := ll.Pop()
	if err != nil {
		t.Errorf("LinkedList: Pop error: %s", err)
	}
	if tNode != node3 {
		t.Errorf("LinkedList: Pop error: it should be node3 but get %v", tNode)
	}
}

func TestLinkedListDeleteByIndex(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	if err := ll.DeleteByIndex(1); err != nil {
		t.Errorf("LinkedList: DeleteByIndex error: %s", err)
	}

	if r := ll.String(); r != "3->1" {
		t.Errorf("LinkedList: String error: it should be 3->1 but get: %s", r)
	}

	if err := ll.DeleteByIndex(999); err == nil {
		t.Errorf("LinkedList: DeleteByIndex error: it should not be nil")
	}

	if err := ll.DeleteByIndex(0); err != nil {
		t.Errorf("LinkedList: DeleteByIndex error: %s", err)
	}
}

func TestLinkedListDeleteByNode(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	ll.DeleteByNode(node1)

	if r := ll.String(); r != "3->2" {
		t.Errorf("LinkedList: String error: it should be 3->2 but get: %s", r)
	}

	if err := ll.DeleteByNode(NewListNode(5)); err == nil {
		t.Errorf("LinkedList: DeleteByNode error: it should not be nil")
	}

}

func TestLinkedListGetFirst(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	if index, node := ll.GetFirst(1); index != 2 || node != node1 {
		t.Errorf("LinkedList: GetFirst error: index and node should be %d and %v, but get %d and %v", 2, node1, index, node)
	}

	if index, node := ll.GetFirst(100); index != -1 || node != nil {
		t.Errorf("LinkedList: GetFirst error: index and node should be %d and %v, but get %d and %v", -1, nil, index, node)
	}
}

func TestLinkedListModify(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	ll.Modify(0, 100)

	if r := ll.String(); r != "100->2->1" {
		t.Errorf("LinkedList: Modify error: it should be 100->2->1 but get %s", r)
	}

	if err := ll.Modify(100, 0); err == nil {
		t.Errorf("LinkedList: Modify error: it should not be nil")
	}
}

func TestLinkedListClear(t *testing.T) {
	var node1 = &ListNode{
		Next:  nil,
		Value: 1,
	}
	var node2 = &ListNode{
		Next:  node1,
		Value: 2,
	}
	var node3 = &ListNode{
		Next:  node2,
		Value: 3,
	}
	ll := NewLinkedList()
	ll.Insert(node1)
	ll.Insert(node2)
	ll.Insert(node3)

	ll.Clear()

	if ll.Head != nil || ll.GetLength() != 0 {
		t.Errorf("LinkedList: Clear error: ll.Head should be nil but is %v, ll.length should be 0 but is %v", ll.Head, ll.GetLength())
	}
}
