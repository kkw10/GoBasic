package main

import "fmt"

func main() {
	var a []int

	a = make([]int, 3)

	a[0] = 100
	a[1] = 200
	a[2] = 300

	fmt.Println(a, len(a), cap(a))

	a = append(a, 400)

	fmt.Println(a, len(a), cap(a))

	var b []int

	b = append(a, 500, 600, 700)

	fmt.Println(b, len(b), cap(b))

	var c []int

	c = append(b, 800)

	fmt.Println(b, len(b), cap(b))

	c[0] = 999

	fmt.Println("//////////////////////////")

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}