package golang

import "testing"

func TestNewHashTable(t *testing.T) {
	h := NewHashTable(64)

	if h.size != 64 {
		t.Fatalf("NewHashTable operation failed: actual %d, expected %d.", h.size, 64)
	}
}

func TestHashTable_Set(t *testing.T) {
	h := NewHashTable(128)
	h.Set("foo", "bar")

	if h.entries != 1 {
		t.Fatalf("actual %d, expected %d. ", h.entries, 1)
	}
}

