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
			i := h.size - 1
			for i > 0 {
				parent := (i - 1)/2
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

func (h *Heap) Remove(n int) int {
	if h.Empty(){
		return -1
	}
	temp := h.priorityQueue[0]
	h.Swap(0, h.size-1)
	h.priorityQueue[h.size - 1] = 0
	h.size--
	if(h.size > 1) {
		i := 0
		for  i<h.size {
			parent := i
			left := (i * 2) + 1
			right := (i * 2) + 2

			// Checks to see if both left & right children exist.
			if(h.priorityQueue[left] != 0 && h.priorityQueue[right] != 0) {
				// Check to see if the parent is greater than either children
				if(h.priorityQueue[parent] > h.priorityQueue[left] || h.priorityQueue[parent] > h.priorityQueue[right]) {
					if(h.priorityQueue[left] > h.priorityQueue[right]) {
						h.Swap(right, parent)
						i = right
					}else {
						h.Swap(left, parent)
						i = left
					}
				}else{
					return temp
				}
			// Checks to see if only the left child exists.
			}else if(h.priorityQueue[left] != 0) {
				if(h.priorityQueue[parent] > h.priorityQueue[left]) {
					h.Swap(left, parent)
					i = left
				}else {
					return temp
				}
			}else{
				i++
			}
		}
	}
	return temp
}

func (h *Heap) Check() bool {
	if(h.Empty()){
		return true
	}
	for i:=0; i<h.size; i++ {
		left := (2*i) + 1
		right := (2*i) + 2

		if(left >= h.size) {
			return true
		}
		if(left != 0) {
			if(h.priorityQueue[i] > h.priorityQueue[left]) {
				return false
			}
		}
		if(right >= h.size) {
			return true
		}
		if(right != 0) {
			if(h.priorityQueue[i] > h.priorityQueue[right]) {
				return false
			}
		}
	}
	return true
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
