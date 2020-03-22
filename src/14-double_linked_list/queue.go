package main

import (
	"fmt"
	"data_structure"
)

func main() {
	queue := data_structure.NewQueue()

	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}

	for !queue.IsEmpty() {
		val := queue.Pop()

		fmt.Printf("%d -> ", val)
	}
}