package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {
	var p Person

	p1 := Person{"tester1", 30}
	p2 := Person{name: "tester2", age: 20}
	p3 := Person{name: "tester3"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "tester4"
	p.age = 25

	fmt.Println(p)

	p.PrintName();
}

// 구조체 = class
// C의 구조체는 오직 속성만을 갖지만 다른 현대의 언어들에서는 구조체가 메서드를 갖는다.
// 메서드를 갖는 구조체를 first class라고 부른다.
// Go의 구조체는 일급객체(first class)이다.
// Go의 구조체의 메서드는 구조체 밖에서 정의된다.