package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := append(a, 3)

	// a의 capacity가 남아있는 상태에서 a에 추가로 할당했으므로 
	// b는 새로운 메모리 주소가 아닌 기존 a의 주소를 가리키게된다.
	fmt.Printf("%p %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)

	// 위와 같은 상황을 방지하려면 먼저 복사를 한후 append를 한다.
	c := make([]int, 2, 4)
	c[0] = 1
	c[1] = 2

	d := make([]int, len(c))

	for i := 0; i < len(c); i++ {
		d[i] = c[i]
	}

	d = append(d, 3)

	fmt.Printf("%p, %p\n", c, d)

	d[0] = 4
	d[1] = 5

	fmt.Println(c)
	fmt.Println(d)
}