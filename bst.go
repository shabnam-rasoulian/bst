package bst

import "fmt"

type Comparable interface {
	Compare(that interface{}) int
}

type node struct {
	data Comparable
	left *node
	right *node
}

func (n *node) insert(data Comparable) bool {
	c := n.data.Compare(data)
	switch {
	case c < 0:
		if n.left == nil {
			n.left = &node {
				data: data,
			}
			return true
		}
		return n.left.insert(data)
	case c == 0:
		return false
	case c > 0:
		if n.right == nil {
			n.right = &node {
				data: data,
			}
			return true
		}
		return n.right.insert(data)
	}
	return false
}

func (n *node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("(%s %v %s)", n.left, n.data, n.right)
}

type Tree struct {
	root *node
}

func (t *Tree) String() string {
	return t.root.String()
}

func (t *Tree) Insert(data Comparable) bool {
	if t.root == nil {
		t.root = &node{
			data: data,
		}
		return true
	}
	return t.root.insert(data)
}
