package main

import "fmt"

func main() {
	var a int
	a = 1

	Increase(a) // 함수에 인자를 전달할때 인자는 참조가 아니라 복사된다.

	fmt.Println(a) // a = 1

	Increase2(&a)

	fmt.Println(a) // a = 2
}

func Increase(x int) {
	x++
}

func Increase2(x *int) {
	*x = *x + 1
}