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
		t.Fatalf("actual %d, want %d", list.head.data, 10)
	}
}

func TestRemove(t *testing.T) {
	list := &List{}
	list.append(10)
	list.append(-10)
	list.append(99999)

	list.remove(99999)
	
	if (list.count != 2) {
		t.Fatalf("remove failed [test by list size]: actual %d, want %d.", list.count, 2)
	}
	if (list.tail.data != -10) {
		t.Fatalf("remove failed [test by tail value]: actual %d, want %d.", list.tail.data, -10)
	}
}

func TestRemoveLast(t *testing.T) {
	list := &List{}
	list.append(10)
	list.append(-10)
	list.append(99999)
	list.removeLast()

	if (list.count != 2) {
		t.Fatalf("removeLast failed [test by list size]: actual %d, want %d.", list.count, 2)
	}

	if (list.tail.data != -10) {
		t.Fatalf("actual %d, want -10", list.tail.data)
	}
}
