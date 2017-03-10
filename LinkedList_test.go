package golang

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := &List{}
	list.append(10)

	if (list.count == 0) {
		t.Fatalf("List is empty.")
	}
	if (list.head.data != 10) {
		t.Fatalf("actual %d, want 10", list.head.data)
	}
}
