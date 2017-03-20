package golang

import (
	"testing"
	"math/rand"
)

func generateTestHeap() *Heap {
	h := Heap{}
	// Seed random number generator
	rand.Seed(42)
	// Random int representing the length of the priority queue
	n := rand.Intn(255)
	// Add elements to the priority queue
	for i:=0; i<n; i++ {
		x := rand.Intn(10000)
		h.Insert(x)
	}

	return &h
}

func TestCreateHeap(t *testing.T) {
	h := Heap{}

	if(!h.Empty()) {
		t.Fatalf("Expecting heap to be empty: actual %d, expected %d", h.size, 0)
	}
}

func TestHeap_Insert(t *testing.T) {
	h := generateTestHeap()

	if(h.Empty()) {
		t.Fatalf("Expecting heap to be non-empty.")
	}
}

func TestHeap_Remove(t *testing.T) {
	h := generateTestHeap()
	h.Insert(1337)

	n := h.Remove(1337)
	if n != 1337 {
		t.Fatalf("Remove operation failed: actual %d, expected %d.", n, 1337)
	}
}

func TestHeap_Check(t *testing.T) {
	h := generateTestHeap()
	if(!h.Check()) {
		t.Fatalf("Heap is out of order.")
	}
}


