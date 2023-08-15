package trees

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type AllowedData interface {
	constraints.Integer
}

type Node[T AllowedData] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T AllowedData](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) find(value T) bool {
	switch {
	case n == nil:
		return false
	case n.value == value:
		return true
	case n.value > value:
		return n.left.find(value)
	case n.value < value:
		return n.right.find(value)
	default:
		return false
	}
}

func (n *Node[T]) insert(value T) {
	switch {
	case n == nil:
		n.right.insert(value)
	case n.value == value:
	case n.value > value:
		if n.left == nil {
			n.left = NewNode[T](value)
			return
		}
		n.left.insert(value)
	case n.right == nil:
		n.right = NewNode[T](value)
	default:
		n.right.insert(value)
	}
}

func (n *Node[T]) print() string {
	if n == nil {
		return ""
	}
	return n.left.print() + " " + strconv.Itoa(int(n.value)) + "" + n.right.print()
}

type BinaryTree[T AllowedData] struct {
	root *Node[T]
}

func NewBinaryTree[T AllowedData]() BinaryTree[T] {
	return BinaryTree[T]{root: nil}
}

func (t *BinaryTree[T]) Insert(value T) {
	if t.root == nil {
		t.root = NewNode[T](value)
		return
	}
	t.root.insert(value)
}

func (t *BinaryTree[T]) Find(value T) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(value)
}

func (t *BinaryTree[T]) Print() string {
	return t.root.print()
}
