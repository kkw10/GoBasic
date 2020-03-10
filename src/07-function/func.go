package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a, b := fun1(2, 3)

	fmt.Println(a, b)

	recursive(10)
}

func fun1(x, y int) (int, int) {
	return y, x
}

func recursive(x int) {
	if x == 0 {
		return
	}

	fmt.Println(x)

	recursive(x-1)
}