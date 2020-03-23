package datastruct

func NewStack() *Stack {
	s := new(Stack)
	s.LinkedList = NewLinkedList()
	s.length = 0
	return s
}

func (s *Stack) Destroy() {
	s.Head = nil
	s.length = 0
}

func (s *Stack) IsEmpty() bool {
	if s.length == 0 && s.Head == nil {
		return true
	}
	return false
}

func (s *Stack) Push(value interface{}) {
	ln := NewListNode(value)
	s.Insert(ln)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.Head
	s.Head = tmp.Next
	s.length--
	return tmp.Value
}

func (s *Stack) GetTop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Head.Value
}
