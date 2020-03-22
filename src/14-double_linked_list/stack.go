package main

import (
	"fmt"
	"data_structure"
)

func main() {
	stack := data_structure.NewStack()

	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	for !stack.IsEmpty() {
		val := stack.Pop()

		fmt.Printf("%d -> ", val)
	}
}