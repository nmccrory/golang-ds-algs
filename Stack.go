package golang

type Stack struct {
	last *StackNode
	count int
}

type StackNode struct {
	data int
	next *StackNode
}

func (s *Stack) push(data int) *Stack {
	n := &StackNode{data: data, next: nil}
	if (s.last != nil) {
		n.next = s.last
	}
	s.last = n
	s.count++
	return s
}

func (s *Stack) pop() (int, *Stack) {
	if (s.last == nil) {
		return -1, s
	}
	res := s.last.data
	s.last = s.last.next
	s.count--
	return res, s
}

func (s *Stack) peek() int {
	if (s.last == nil) {
		return -1
	}
	return s.last.data
}

func (s *Stack) size() int {
	return s.count
}