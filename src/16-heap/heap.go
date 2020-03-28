package main

import (
	"fmt"
	"data_structure"
)

func main() {
	h := &data_structure.Heap{}

	h.Push(2)
	h.Push(5)
	h.Push(1)
	h.Push(8)
	h.Push(3)
	h.Push(4)
	h.Push(1)
	h.Push(3)
	h.Push(7)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}