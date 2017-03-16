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

func minValue(n *TreeNode) int {
	if n.left == nil {
		return n.data
	}
	return minValue(n.right)
}

func link(n *TreeNode) {
	if n.parent.left == n {
		if n.left != nil {
			n.parent.left = n.left
		}else{
			n.parent.left = n.right
		}
	} else if n.parent.right == n {
		if n.left != nil {
			n.parent.right = n.left
			n.parent.right = n.right
		}
	}
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

func Remove(parent *TreeNode, n *TreeNode, data int) bool {
	switch {
	case n.data == data:
		if n.left != nil && n.right != nil {
			n.data = minValue(n.right)
			return Remove(n.right, n, n.data)
		}
		link(n); return true
	case n.data < data :
		if n.right == nil {
			return false
		}
		return Remove(n, n.right, n.data)
	case n.data > data :
		if n.left == nil {
			return false
		}
		return Remove(n, n.left, n.data)
	}
	return false
}

func (t *Tree) Remove(data int) bool {
	if t.isEmpty() {
		return false
	}
	exists, _ := t.Find(data)
	if !exists {
		return false
	}
	if t.root.data == data {
		temp := &TreeNode{0, nil, t.root, nil}
		success := Remove(temp, t.root, data)
		t.root = temp.left
		return success
	}
	return Remove(t.root, t.root.left, data) || Remove(t.root, t.root.right, data)
}

func (n *TreeNode) Find(data int) (bool, *TreeNode) {
	if (n.data == data) {
		return true, n
	}
	if (data > n.data) {
		return n.right.Find(data)
	} else if(data < n.data){
		return n.left.Find(data)
	}
	return false, nil
}

func (t *Tree) Find(data int) (bool, *TreeNode) {
	if (t.isEmpty()) {
		return false, nil
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