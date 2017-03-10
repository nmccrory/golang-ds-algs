// I know it's called 'GolangAlgorithms' but this is a data structure. :P
package golang

type Node struct {
	data int
	next, prev *Node
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