package main

import "fmt"

func main() {
	var a int
	var b int
	var p *int // 포인터는 타입이다.

	a = 3
	b = 2
	p = &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	 p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
}

// C, C++은 포인터를 사용한다.
// 포인터는 메모리 주소를 나타내고 이를 사용해서 연산이나 형변환등을 할 수 있다.
// 포인터를 이용한 연산은 어렵다.

// 포인터는 "이해하기 어렵다", "메모리에 직접 접근해서 위험할 수 있다"는 이유로인해 대부분의 현대 언어(C#, JAVA..)에서는 명시적으로 사용되고 있지 않다.( 내부적으로는 사용됨. )
// 하지만 포인터를 사용하지 않기 때문에 발생하는 문제가 있다.
// 1. 내부 기능을 명확하기 이해하기 어렵다.
// 2. 타입이 pointer 타입인지 value 타입인지 프로그래머가 일일이 알고 있어야 한다.
// (이를 위해서 C#은 타입을 reference, value로 나눈다.)

// 종합해보면 pointer를 사용해서 또는 사용하지 않아서 발생하는 문제는 다음과 같다.
// C, C++ => 포인터를 활용해서 연산이 가능하지만 매우 어려워서 프로그래머가 이해하기 쉽지 않다.
// JAVA, C# => 포인터를 내부적으로 사용하고 있음에도 명시적으로 사용되지 않아서 프로그램을 명확하 이해하기 어렵게 한다.

// GO는 위와 같은 상황을 절충해서 다음과 같이 포인터를 사용한다.
// 포인터를 타입의 형태로 사용할 수 있다.
// 포인터를 활용한 연산은 할 수 없다.

// [포인터 정리]
// 포인터 vs 변수
// 포인터 : 변수(메모리 주소)를 가리킨다. 
// 변수 : 메모리 주소 + 메모리에 저장된 값