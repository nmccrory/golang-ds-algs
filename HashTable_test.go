package golang

import "testing"

func TestHashTable_PrintHash(t *testing.T) {
	h := HashTable{}
	h.PrintHash("hello")
}
