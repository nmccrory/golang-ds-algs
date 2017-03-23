package golang

import "testing"
import "fmt"

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

func TestHashTable_Get(t *testing.T) {
	h := NewHashTable(128)
	h.Set("foo", "bar")
	h.Set("hello", "World")

	val := h.Get("foo")
	if val != "bar" {
		t.Fatalf("actual %d, expected %d.", val, "bar")
	}
}

func TestHashTable_Delete(t *testing.T) {
	h := NewHashTable(128)
	h.Set("big", "foot")
	h.Set("John", "Legend")
	h.Set("foo", "bar")
	h.Set("United States", "OF AMERICA")

	h.Delete("foo")

	if h.entries != 3 {
		t.Fatalf("Delete failed. actual size %d, expected %d", h.entries, 3)
	}
}