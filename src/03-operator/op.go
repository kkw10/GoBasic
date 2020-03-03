package main

import "fmt"

func main() {
	// 변수의 선언과 할당을 동시에 함
	a := 1;
	b := 2;

	// 비트 연산자 사용해보기
	fmt.Printf("a & b %v\n", a&b) // AND
	fmt.Printf("a | b %v\n", a|b) // OR
	fmt.Println("result", a^b) // XOR

	// XOR의 경우 단항 연산자로 쓰이면 NOT이 된다.

	var c bool = 3 > 4

	fmt.Println(c)
}