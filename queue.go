package datastruct

func NewQueue() *Queue {
	q := new(Queue)
	q.LinkedList = NewLinkedList()
	q.length = 0
	return q
}

func (q *Queue) IsEmpty() bool {
	if q.length == 0 && q.Head == nil {
		return true
	}
	return false
}

func (q *Queue) De() interface{} {
	if q.IsEmpty() {
		return nil
	}
	tmp := q.Head
	q.Head = tmp.Next
	q.length--
	return tmp.Value
}

func (q *Queue) En(vaule interface{}) {
	node := NewListNode(vaule)
	q.InsertAt(node, q.length)
}
