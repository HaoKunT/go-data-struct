package dataStruct

import (
	"testing"
)

func TestStackIsEmpty(t *testing.T) {
	var stack = NewStack()
	if !stack.IsEmpty() {
		t.Errorf("Stack: IsEmpty error: it should be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("Stack: IsEmpty error: it should not be empty")
	}
}

func TestStackPush(t *testing.T) {
	var stack = NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.GetLength() != 3 {
		t.Errorf("Stack: Push error: the stack length should be 3 but get length: %d", stack.GetLength())
	}
}

func TestStackPop(t *testing.T) {
	var stack = NewStack()

	if pop := stack.Pop(); pop != nil {
		t.Errorf("Stack: Pop error: it should be nil, but get %v", pop)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if pop := stack.Pop(); pop != 3 {
		t.Errorf("Stack: Pop error: the node should be %d, but is %d", 3, pop)
	}
	if stack.GetLength() != 2 {
		t.Errorf("Stack: Pop error: the stack length should be %d, but is %d", 2, stack.GetLength())
	}
}

func TestStackGetTop(t *testing.T) {
	var stack = NewStack()

	if top := stack.GetTop(); top != nil {
		t.Errorf("Stack: GetTop error: it should be nil, but get %v", top)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if top := stack.GetTop(); top != 3 {
		t.Errorf("Stack: GetTop error: the node should be %d, but is %d", 3, top)
	}
	if stack.GetLength() != 3 {
		t.Errorf("Stack: GetTop error: the stack length should be %d, but is %d", 3, stack.GetLength())
	}
}
