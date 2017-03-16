package golang

import (
	"fmt"
	"time"
)

type Heap struct {
	priorityQueue [255]int
	size int
}

type HeapError struct {
	When time.Time
	What string
}

func (e HeapError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// Less(a, b) compares a to b and returns true if a is smaller than b.
func Less(a int, b int) bool {
	if a < b {
		return true
	}
	return false
}

func (h *Heap) Swap(a int, b int) {
	temp := h.priorityQueue[a]
	h.priorityQueue[a] = h.priorityQueue[b]
	h.priorityQueue[b] = temp
}

func (h *Heap) Empty() bool {
	if h.size > 0 {
		return false
	}
	return true
}

func (h *Heap) Min() int {
	return h.priorityQueue[0]
}

func (h *Heap) Insert(n int) error {
	if(h.size >= 255) {
		return HeapError{time.Now(), "You have exceeded the size of the Priority Queue."}
	}else{
		h.priorityQueue[h.size] = n
		h.size++
		if(h.size > 1) {
			for i := h.size - 1; i > 0; i++ {
				parent := (i -1)/2
				if(Less(h.priorityQueue[i], h.priorityQueue[parent])) {
					h.Swap(i, parent)
					i = parent
				}else{
					break
				}
			}
		}
		return nil
	}
}

func (h *Heap) PrintPriorityQueue() {
	fmt.Printf("\n Priority Queue Print: ")
	for i:=0; i<h.size; i++ {
		if i == 0 {
			fmt.Printf("\n[")
		}
		fmt.Printf(" %d ", h.priorityQueue[i])
		if i == h.size - 1 {
			fmt.Printf("]\n")
		}
	}
}
