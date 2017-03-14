package golang

import "testing"

func TestNewTree(t *testing.T) {
	tree := Tree{root: nil, size: 0}

	if tree.root != nil || tree.size != 0 {
		t.Fatalf("Tree could not be initialized.")
	}
}

func TestTreeSize(t *testing.T) {
	tree := Tree{root: nil, size: 0}

	if tree.Size() != 0 {
		t.Fatalf("Size value is not correct: actual %d, expected %d", tree.Size(), 0)
	}
}

func TestInsertTreeNode(t *testing.T) {
	tree := Tree{root: nil, size: 0}
	tree.insert(10)
	tree.insert(-10)
	tree.insert(9999)
	tree.insert(1)
	tree.insert(2)

	if tree.Size() != 5 {
		t.Fatalf("Size is not correct: actual %d, expected %d", tree.Size(), 5)
	}
}

func TestFindTreeNode(t *testing.T) {
	tree := Tree{root: nil, size: 0}
	tree.insert(10)
	tree.insert(-10)
	tree.insert(9999)
	tree.insert(1)
	tree.insert(404)
	tree.insert(2)
	result := tree.find(404)

	if result.data != 404 {
		t.Fatalf("Find operation failed. actual %d, expected %d", result.data, 404)
	}
}