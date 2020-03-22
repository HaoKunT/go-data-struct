package dataStruct

import "fmt"

func NewBinaryTree() *BinaryTree {
	return new(BinaryTree)
}

func NewBinaryTreeFromNode(node *BinaryTreeNode) *BinaryTree {
	bt := NewBinaryTree()
	bt.Root = node
	return bt
}

func NewBinaryTreeNode(value interface{}) *BinaryTreeNode {
	btn := new(BinaryTreeNode)
	btn.Value = value
	return btn
}

func (bt *BinaryTree) InsertLeftChild(lt *BinaryTree) {
	bt.Root.Left = lt.Root
}

func (bt *BinaryTree) InsertRightChild(rt *BinaryTree) {
	bt.Root.Right = rt.Root
}

func (bt *BinaryTree) Left() *BinaryTree {
	return NewBinaryTreeFromNode(bt.Root.Left)
}

func (bt *BinaryTree) Right() *BinaryTree {
	return NewBinaryTreeFromNode(bt.Root.Right)
}

func (bt *BinaryTree) LeftNode() *BinaryTreeNode {
	return bt.Root.Left
}

func (bt *BinaryTree) RightNode() *BinaryTreeNode {
	return bt.Root.Right
}

func (bt *BinaryTree) DeleteLeft() {
	bt.Root.Left = nil
}

func (bt *BinaryTree) DeleteRight() {
	bt.Root.Right = nil
}

func (bt *BinaryTree) PreOrderTraverseRecursive(f BinaryTreeHandleFunc, e BinaryTreeErrorHandleFunc) {
	tNode := bt.Root
	if tNode == nil {
		return
	}
	err := f(tNode)
	if err != nil {
		e(err)
	}
	l := NewBinaryTreeFromNode(tNode.Left)
	l.PreOrderTraverseRecursive(f, e)
	r := NewBinaryTreeFromNode(tNode.Right)
	r.PreOrderTraverseRecursive(f, e)
}

// func (bt *BinaryTree) PreOrderTraverseStack(f BinaryTreeHandleFunc, e BinaryTreeErrorHandleFunc) {
// 	// now we do not have stack data struct
// 	panic("Sorry we do not have stack now")
// }

func (bt *BinaryTree) InOrderTraverseRecursive(f BinaryTreeHandleFunc, e BinaryTreeErrorHandleFunc) {
	tNode := bt.Root
	if tNode == nil {
		return
	}
	l := NewBinaryTreeFromNode(tNode.Left)
	l.InOrderTraverseRecursive(f, e)
	err := f(tNode)
	if err != nil {
		e(err)
	}
	r := NewBinaryTreeFromNode(tNode.Right)
	r.InOrderTraverseRecursive(f, e)
}

func (bt *BinaryTree) PostOrderTraverseRecursive(f BinaryTreeHandleFunc, e BinaryTreeErrorHandleFunc) {
	tNode := bt.Root
	if tNode == nil {
		return
	}
	l := NewBinaryTreeFromNode(tNode.Left)
	l.PostOrderTraverseRecursive(f, e)
	r := NewBinaryTreeFromNode(tNode.Right)
	r.PostOrderTraverseRecursive(f, e)
	err := f(tNode)
	if err != nil {
		e(err)
	}
}

func (bt *BinaryTree) LeverTraverse(f BinaryTreeHandleFunc, e BinaryTreeErrorHandleFunc) {
	var q = make([]*BinaryTreeNode, 0)
	q = append(q, bt.Root)
	for len(q) > 0 {
		err := f(q[0])
		if err != nil {
			e(err)
		}
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		q = q[1:]
	}
}

func (bt *BinaryTree) Deep() int {
	if bt.Root == nil {
		return 0
	}
	l := NewBinaryTreeFromNode(bt.Root.Left)
	nLeft := l.Deep()
	r := NewBinaryTreeFromNode(bt.Root.Right)
	nRight := r.Deep()
	if nLeft > nRight {
		return nLeft + 1
	}
	return nRight + 1
}

func (btn *BinaryTreeNode) String() string {
	return fmt.Sprintf("%v", btn.Value)
}