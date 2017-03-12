package golang

import "testing"


func TestPush(t *testing.T) {
	s := Stack{last: nil, count: 0}
	s.push(10)
	s.push(9999)

	if (s.last.data != 9999) {
		t.Fatalf("actual %d, expected %d", s.last.data, 9999)
	}
}

func TestSize(t *testing.T) {
	s := Stack{last: nil, count: 0}

	if (s.size() != 0) {
		t.Fatalf("Size is incorrect. Actual %d, expected %d", s.size(), 0)
	}
	s.push(10)
	s.push(9999)
	if (s.size() != 2) {
		t.Fatalf("Size is incorrect. Actual %d, expected %d", s.size(), 2)
	}
}

func TestPop(t *testing.T) {
	s := Stack{last: nil, count: 0}
	s.push(10)
	s.push(9999)
	s.push(-10)
	s.push(0)

	val, s2 := s.pop()

	if (val != 0) {
		t.Fatalf("actual %d, expected %d", val, 0)
	}

	if (s2.count != 3) {
		t.Fatalf("actual size %d, expected size %d", s2.count, 3)
	}
}

func TestPeek(t *testing.T) {
	s := Stack{last: nil, count: 0}
	s.push(10)
	s.push(9999)
	s.push(5678)

	if (s.peek() != 5678) {
		t.Fatalf("actual %d, expected %d", s.peek(), 5678)
	}
}
