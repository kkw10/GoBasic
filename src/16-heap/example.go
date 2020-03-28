package main

import (
	"fmt"
	"data_structure"
)

func main() {
	// Quest 1 
	h := &data_structure.MinHeap{}

	nums1 := []int{-1, 3, -1, 5, 4}
	size1 := 2

	for i := 0; i < len(nums1); i++ {
		h.Push(nums1[i])

		if h.Count() > size1 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())

	// Quest 2
	h2 := &data_structure.MinHeap{}

	nums2 := []int{2, 4, -2, -3, 8}
	size2 := 1

	for i := 0; i < len(nums2); i++ {
		h2.Push(nums2[i])

		if h2.Count() > size2 {
			h2.Pop()
		}
	}

	fmt.Println(h2.Pop())
	
	// Quest 3
	h3 := &data_structure.MinHeap{}

	nums3 := []int{-5, -3, 1}
	size3 := 3

	for i := 0; i < len(nums3); i++ {
		h3.Push(nums3[i])

		if h3.Count() > size3 {
			h3.Pop()
		}
	}

	fmt.Println(h3.Pop())
}