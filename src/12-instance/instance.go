package main

import "fmt"

type Student struct {
	name string
	age int
	grade int
}

// Function (Procedure) ...
func SetName(t Student, newName string) {
	t.name = newName
}

// OOP...
func (t *Student) SetName(newName string) {
	t.name = newName
}

func main() {
	// Value 타입의 경우
	a := Student{"Jonathan", 40, 10}
	b := a

	b.age = 30
	SetName(b, "Rick") // b"를" SetName 한다.

	fmt.Println(a)
	fmt.Println(b)
	
	fmt.Println("////////////////////")

	// Reference 타입의 경우
	var d *Student
	c := Student{"Ethan", 40, 10}
	d = &c

	d.age = 30
	d.SetName("Rick") // d"가" SetName 한다.
	// 즉 b(= instance)가 주어가 된다.
	// 이는 b가 프로그래밍 적으로 생명주기를 갖는다는 의미이다.

	fmt.Println(c)
	fmt.Println(d)
}