package golang

import (
	"testing"
	"fmt"
)

func generateTestTree() *Tree {
	tree := Tree{root: nil, size: 0}
	tree.Insert(10)
	tree.Insert(-10)
	tree.Insert(9999)
	tree.Insert(1)
	tree.Insert(404)
	tree.Insert(2)
	tree.Insert(25)
	return &tree
}

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
	tree := generateTestTree()

	if tree.Size() != 7 {
		t.Fatalf("Size is not correct: actual %d, expected %d", tree.Size(), 7)
	}
}

func TestFindTreeNode(t *testing.T) {
	tree := generateTestTree()
	found, n := tree.Find(404)
	if found {
		if n.data != 404 {
			t.Fatalf("Find operation failed. actual %d, expected %d", n.data, 404)
		}
	}else{
		t.Fatalf("Node does not exist.")
	}
}

func TestPreorderPrint(t *testing.T) {
	tree := generateTestTree()
	fmt.Printf("*Preorder Traversal*\n")
	tree.PreorderPrint()
	fmt.Printf("\n\n")
}

func TestPostorderPrint(t *testing.T) {
	tree := generateTestTree()
	fmt.Printf("*Postorder Traversal*\n")
	tree.PostorderPrint()
	fmt.Printf("\n\n")
}

func TestInorderPrint(t *testing.T) {
	tree := generateTestTree()
	fmt.Printf("*Inorder Traversal*\n")
	tree.InorderPrint()
	fmt.Printf("\n\n")
}

func TestTreeRemove(t *testing.T) {
	//tree := generateTestTree()

	//if !tree.Remove(404) {
	//	t.Fatalf("Remove Node operation failed. [Node %d]", 404)
	//}
	//tree.PreorderPrint()
}