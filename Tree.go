package golang

import "fmt"

type Tree struct {
	root *TreeNode
	size int
}

type TreeNode struct {
	data int
	parent *TreeNode
	left *TreeNode
	right *TreeNode
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Root() *TreeNode {
	return t.root
}

func (t *Tree) isEmpty() bool {
	if t.Size() == 0 {
		return true
	}
	return false
}

func (root *TreeNode) Insert(n *TreeNode) {
	if n.data > root.data {
		if root.right == nil {
			root.right = n
			n.parent = root
		}else{
			root.right.Insert(n)
		}
	} else if (n.data < root.data){
		if root.left == nil {
			root.left = n
			n.parent = root
		}else{
			root.left.Insert(n)
		}
	}
}

func (t *Tree) Insert(data int) {
	n := &TreeNode{data: data, parent: nil, left: nil, right: nil}

	if t.root == nil {
		t.root = n
	}
	t.size++
	t.root.Insert(n)
}

func Remove(parent *TreeNode, data int) {
	// TODO: finish making remove function
}

func (t *Tree) Remove(data int) {
	if t.isEmpty() {
		return
	}
	t.Find(data)
	Remove(t.Root(), data)
}

func (n *TreeNode) Find(data int) *TreeNode {
	if (n.data == data) {
		return n
	}
	if (data > n.data) {
		return n.right.Find(data)
	} else{
		return n.left.Find(data)
	}
}

func (t *Tree) Find(data int) *TreeNode {
	if (t.isEmpty()) {
		return nil
	}
	return t.root.Find(data)
}

func PreorderPrint(n *TreeNode) {
	if n != nil {
		fmt.Printf("[%d]", n.data)
		PreorderPrint(n.left)
		PreorderPrint(n.right)
	}

}

func (t *Tree) PreorderPrint() {
	if t.isEmpty() {
		fmt.Print("Tree is empty.")
		return
	}
	PreorderPrint(t.root)
}

func PostorderPrint(n *TreeNode) {
	if n != nil {
		PostorderPrint(n.left)
		PostorderPrint(n.right)
		fmt.Printf("[%d]", n.data)
	}
}

func (t *Tree) PostorderPrint() {
	if t.isEmpty() {
		fmt.Print("Tree is empty.")
		return
	}
	PostorderPrint(t.root)
}

func InorderPrint(n *TreeNode) {
	if n != nil {
		InorderPrint(n.left)
		fmt.Printf("[%d]", n.data)
		InorderPrint(n.right)
	}
}

func (t *Tree) InorderPrint() {
	if t.isEmpty() {
		fmt.Print("Tree is empty.")
		return
	}
	InorderPrint(t.root)
}