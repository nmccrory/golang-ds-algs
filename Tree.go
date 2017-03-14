package golang

type Tree struct {
	root *TreeNode
	size int
}

type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
}

func (root *TreeNode) insert(n *TreeNode) {
	if n.data > root.data {
		if root.right == nil {
			root.right = n
		}else{
			root.right.insert(n)
		}
	} else if (n.data < root.data){
		if root.left == nil {
			root.left = n
		}else{
			root.left.insert(n)
		}
	}
}

func (t *Tree) insert(data int) {
	n := &TreeNode{data: data, left: nil, right: nil}

	if t.root == nil {
		t.root = n
	}
	t.size++
	t.root.insert(n)
}

func (t *Tree) remove(data int) {
	if t.isEmpty() {
		return
	}
	if t.Size() == 1 && t.root.data != data {
		return
	}

}

func (n *TreeNode) find(data int) *TreeNode {
	if (n.data == data) {
		return n
	}
	if (data > n.data) {
		return n.right.find(data)
	} else{
		return n.left.find(data)
	}
}

func (t *Tree) find(data int) *TreeNode {
	if (t.isEmpty()) {
		return nil
	}
	return t.root.find(data)
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