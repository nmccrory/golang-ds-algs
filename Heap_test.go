package golang

import (
	"testing"
)

func TestCreateHeap(t *testing.T) {
	h := Heap{}
	h.Insert(11)
	h.Insert(10)
	h.Insert(12)
	h.Insert(14)
	h.PrintPriorityQueue()
}
