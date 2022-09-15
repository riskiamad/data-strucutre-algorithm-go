package main

import "fmt"

// heaps : tree based data structure
type MaxHeap struct {
	array []int
}

// Insert : adds an element into the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIdx := len(h.array) - 1

	if len(h.array) == 0 {
		println("cannot extract array if length 0")
		return -1
	}

	// take the last index and put in the root
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(idx int) {

	lastIdx := len(h.array) - 1
	l, r := left(idx), right(idx)
	chidlToCompare := 0

	// loop while index has at least one child
	for l <= lastIdx {
		//when left child is the only one child
		if l == lastIdx {
			chidlToCompare = l
		} else if h.array[l] > h.array[r] {
			chidlToCompare = l
		} else {
			chidlToCompare = r
		}

		if h.array[idx] < h.array[chidlToCompare] {
			h.swap(idx, chidlToCompare)
			idx = chidlToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

// get the parent index
func parent(idx int) int {
	return (idx - 1) / 2
}

// get the left child index
func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}

//swap keys in array
func (h *MaxHeap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func main() {
	m := &MaxHeap{}
	println(m)

	buildHeap := []int{10, 20, 30, 7, 9, 13, 17, 21, 35, 40}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	println("Extract:")
	for i := 0; i < len(buildHeap); i++ {
		m.Extract()
		fmt.Println(m)
	}
}
