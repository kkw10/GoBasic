package main

import "fmt"

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a) - 1], a[len(a) - 1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	// slice를 slice해보자.
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	c := a[4:]
	d := a[:4]

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// slice는 배열을 가리키는 포인터이다.
	// slice를 slice할때 실제 배열을 잘라서 가져오는 것이 아니라 해당 부부만 가리키는 포인터를 가져오는 것이다.
	b[0] = 1
	b[1] = 2

	fmt.Println(a)

	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println()
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		var front int
		a, front = RemoveFront(a)
		fmt.Printf("%d, ", front)
	}
	fmt.Println()
	fmt.Println(a)
}