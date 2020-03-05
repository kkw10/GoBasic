package main

import "fmt"

func main() {
	i := 0

	for (i < 10) {
		fmt.Println(i)
		i++
	}

	fmt.Printf("최종 i값은 : %d\n", i)

	// 변수의 스코프를 주의한다.
	var j int

	// 반복문의 조건문에서 만약에 j := 0;을 사용했다면 어떻게 될까?
	// 반복문에서 사용한 j 변수와 앞서 선언한 j 변수는 이름은 같지만 스코프가 다른 변수이다.
	// 따라서 아래 마지막 프린트 함수에서 j의 값은 기본 값인 0이 된다. ( go에서는 변수에 값을 할당하지 않으면 기본 값을 할당한다. )
	for j = 0; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("최종 j값은 :", j)
}