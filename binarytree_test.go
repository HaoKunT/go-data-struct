package datastruct

import (
	"bytes"
	"testing"
)

func TestNewBinaryTreeFromNode(t *testing.T) {
	var node1 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 1,
	}
	bt := NewBinaryTreeFromNode(node1)
	if bt.Root != node1 {
		t.Errorf("BinaryTree: NewBinaryTreeFromNode error: root should be %v, but is %v", node1, bt.Root)
	}
}

func TestBinaryTreeInsertLeftChild(t *testing.T) {
	var node1 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 1,
	}
	var node2 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 2,
	}
	var node3 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 3,
	}
	t1 := NewBinaryTreeFromNode(node1)
	t2 := NewBinaryTreeFromNode(node2)
	t3 := NewBinaryTreeFromNode(node3)
	t1.InsertLeftChild(t2)
	t2.InsertLeftChild(t3)

	if t1.Root.Left != node2 || t1.Root.Left.Left != node3 {
		t.Errorf("BinaryTree: InsertLeftChild error, it should be \" t1.Root.Left == node2 || t1.Root.Left.Left == node3\"")
	}
}

func TestBinaryTreeInsertRightChild(t *testing.T) {
	var node1 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 1,
	}
	var node2 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 2,
	}
	var node3 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 3,
	}
	t1 := NewBinaryTreeFromNode(node1)
	t2 := NewBinaryTreeFromNode(node2)
	t3 := NewBinaryTreeFromNode(node3)
	t1.InsertRightChild(t2)
	t2.InsertRightChild(t3)

	if t1.Root.Right != node2 || t1.Root.Right.Right != node3 {
		t.Errorf("BinaryTree: InsertLeftChild error, it should be \" t1.Root.Right == node2 || t1.Root.Right.Right == node3\"")
	}
}

func TestBinaryTreeLeft(t *testing.T) {
	var node1 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 1,
	}
	var node2 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 2,
	}
	var node3 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 3,
	}
	t1 := NewBinaryTreeFromNode(node1)
	t2 := NewBinaryTreeFromNode(node2)
	t3 := NewBinaryTreeFromNode(node3)
	t1.InsertLeftChild(t2)
	t2.InsertLeftChild(t3)

	if t1.Left().Root != t2.Root || t2.Left().Root != t3.Root {
		t.Errorf("BinaryTree: Left error, it should be \" t1.Left().Root == t2.Root && t2.Left().Root == t3.Root\"")
	}

	if t1.LeftNode() != t2.Root || t2.LeftNode() != t3.Root {
		t.Errorf("BinaryTree: LeftNode error, it should be t1.LeftNode() != t2.Root || t2.LeftNode() != t3.Root")
	}

	if t1.DeleteLeft(); t1.LeftNode() != nil {
		t.Errorf("BinaryTree: DeleteLeft error, it should be nil but get %v", t1.LeftNode())
	}
}

func TestBinaryTreeRight(t *testing.T) {
	var node1 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 1,
	}
	var node2 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 2,
	}
	var node3 = &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: 3,
	}
	t1 := NewBinaryTreeFromNode(node1)
	t2 := NewBinaryTreeFromNode(node2)
	t3 := NewBinaryTreeFromNode(node3)
	t1.InsertRightChild(t2)
	t2.InsertRightChild(t3)

	if t1.Right().Root != t2.Root || t2.Right().Root != t3.Root {
		t.Errorf("BinaryTree: Right error, it should be \" t1.Right().Root == t2.Root && t2.Right().Root == t3.Root\"")
	}

	if t1.RightNode() != t2.Root || t2.RightNode() != t3.Root {
		t.Errorf("BinaryTree: RightNode error, it should be t1.RightNode() != t2.Root || t2.RightNode() != t3.Root")
	}

	if t1.DeleteRight(); t1.RightNode() != nil {
		t.Errorf("BinaryTree: DeleteRight error, it should be nil but get %v", t1.RightNode())
	}
}

func TestBinaryTreePreOrderTraverseRecursive(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	var buf bytes.Buffer

	tree.PreOrderTraverseRecursive(func(node *BinaryTreeNode) error {
		buf.WriteString(node.String())
		return nil
	}, func(err error) {
		t.Error(err)
	})

	if buf.String() != "12453" {
		t.Errorf("BinaryTree: PreOrderTraverseRecursive error: buf.String() should be '12453' but is '%s'", buf.String())
	}
}

func TestBinaryTreeInOrderTraverseRecursive(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	var buf bytes.Buffer

	tree.InOrderTraverseRecursive(func(node *BinaryTreeNode) error {
		buf.WriteString(node.String())
		return nil
	}, func(err error) {
		t.Error(err)
	})

	if buf.String() != "42513" {
		t.Errorf("BinaryTree: InOrderTraverseRecursive error: buf.String() should be '42513' but is '%s'", buf.String())
	}
}

func TestBinaryTreePostOrderTraverseRecursive(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	var buf bytes.Buffer

	tree.PostOrderTraverseRecursive(func(node *BinaryTreeNode) error {
		buf.WriteString(node.String())
		return nil
	}, func(err error) {
		t.Error(err)
	})

	if buf.String() != "45231" {
		t.Errorf("BinaryTree: PostOrderTraverseRecursive error: buf.String() should be '45231' but is '%s'", buf.String())
	}
}

func TestBinaryTreeLeverTraverse(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	var buf bytes.Buffer

	tree.LeverTraverse(func(node *BinaryTreeNode) error {
		buf.WriteString(node.String())
		return nil
	}, func(err error) {
		t.Error(err)
	})

	if buf.String() != "12345" {
		t.Errorf("BinaryTree: LeverTraverse error: buf.String() should be '12345' but is '%s'", buf.String())
	}
}

func TestBinaryTreeDeep(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	if tree.Deep() != 3 {
		t.Errorf("BinaryTree: Deep error: it should be 3 but get %d", tree.Deep())
	}
}

//TestBinaryTreePreOrderTraverseStack
func TestBinaryTreePreOrderTraverseStack(t *testing.T) {
	tree := NewBinaryTreeFromNode(NewBinaryTreeNode(1))
	tree.Root.Left = NewBinaryTreeNode(2)
	tree.Root.Right = NewBinaryTreeNode(3)
	tree.Root.Left.Left = NewBinaryTreeNode(4)
	tree.Root.Left.Right = NewBinaryTreeNode(5)

	var buf bytes.Buffer

	tree.PreOrderTraverseStack(func(node *BinaryTreeNode) error {
		buf.WriteString(node.String())
		return nil
	}, func(err error) {
		t.Error(err)
	})

	if buf.String() != "12453" {
		t.Errorf("BinaryTree: TestBinaryTreePreOrderTraverseStack error: buf.String() should be '12453' but is '%s'", buf.String())
	}
}
