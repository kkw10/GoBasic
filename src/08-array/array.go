package main

import "fmt"

func main() {
	// 배열 기본
	var A [10]int

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}

	fmt.Println(A)

	// 바이트 문자열
	S1 := "Hello 월드"

	for i := 0; i < len(S1); i++ {
		fmt.Print(S1[i], " = ")
		fmt.Print(string(S1[i]), ", ")
	}

	fmt.Println()

	// 룬 문자열
	S2 := []rune(S1)

	for i := 0; i < len(S2); i++ {
		fmt.Print(S2[i], " = ")
		fmt.Print(string(S2[i]), ", ")
	}
}

// 배열은 메모리이다.
// 배열의 길이는 원소의 사이즈 * 배열의 길이이다.
// 문자열 또한 배열이다.
// 영문자는 1byte, 한글은 3byte이다.
// 문자열은 바이트 배열 또는 룬 배열로 나타낼 수 있다.
// rune 배열은 utf8을 나타내는 타입이다. 따라서 가변적이다.