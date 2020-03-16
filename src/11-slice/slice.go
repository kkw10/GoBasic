package main

import "fmt"

func main() {
	var a []int

	fmt.Printf("len(a) = %d\n", len(a)) // len => 배열의 길이
	fmt.Printf("cap(a) = %d\n", cap(a)) // cap => 배열을 위해서 확보해 놓은 공간

	a = append(a, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	b := []int{1, 2, 3, 4, 5}

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	c := make([]int, 0, 8)

	fmt.Printf("len(c) = %d\n", len(c))
	fmt.Printf("cap(c) = %d\n", cap(c))
}
