package main

import "fmt"

func main() {
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
		fmt.Println(last)
	}

	fmt.Println("///////////////////////")

	queue := []int{}

	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var first int
		first, queue = queue[0], queue[1:]
		fmt.Println(first)
	}
}