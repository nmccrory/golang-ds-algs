// I know it's called 'GolangAlgorithms' but this is a data structure. :P
package golang

type Node struct {
	data int
	next, prev *Node
}

func (n *Node) hasNext() bool {
	if (n.next != nil) {
		return true
	}
	return false
}

func (n *Node) hasPrev() bool {
	if (n.prev != nil) {
		return true
	}
	return false
}

type List struct {
	head *Node
	tail *Node
	count int
}

func (l *List) append(data int) *List {
	n := &Node{data: data}
	if (l.head == nil) {
		l.head = n
		l.count = 1
	}else{
		n.prev = l.tail
		l.tail.next = n
		l.count++
	}
	l.tail = n

	return l
}

func (l *List) find(data int) *Node {
	found := false
	var result *Node = nil
	for n := l.head; n != nil && !found; n = n.next {
		if (n.data == data) {
			found = true
			result = n
		}
	}
	return result
}

func (l *List) remove(data int) *List {
	found := false
	if (l.head == nil) {
		return l
	}

	for n := l.head; n != nil && !found; n = n.next {
		if (n.data == data) {
			if (n == l.head) {
				n.next.prev = nil
				l.head = n.next
			}else if (n == l.tail) {
				n.prev.next = nil
				l.tail = n.prev
			}else {
				n.prev.next = n.next
				n.next.prev = n.prev
			}
			found = true
			l.count--
		}
	}
	return l
}

func (l *List) removeLast() *List {
	if (l.tail == nil) {
		return l
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	l.count--
	return l
}

func (l *List) isEmpty() bool {
	if (l.head == nil) {
		return true
	}
	return false
}