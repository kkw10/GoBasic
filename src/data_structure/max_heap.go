package data_structure

import "fmt"

type MaxHeap struct {
	list []int
}

func (h *MaxHeap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx < 0 {
			break;
		}

		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	first := h.list[0]
	last := h.list[len(h.list) - 1]
	h.list = h.list[:len(h.list) - 1]

	if len(h.list) == 0 {
		return first
	}

	h.list[0] = last

	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx * 2 + 1

		if leftIdx >= len(h.list) {
			break
		}

		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx * 2 + 2

		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}

	return first
}

func (h *MaxHeap) Print() {
	fmt.Println(h.list)
}