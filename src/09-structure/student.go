package main

import "fmt"

type Student struct {
	name string
	class int
	point Point
}

type Point struct {
	name string
	grade string
}

func (s Student) ViewPoint() {
	fmt.Println(s.point)
}

func (s Student) InputPoint(name string, grade string) {
	s.point.name = name
	s.point.grade = grade
} 

func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.point.name = "수학"
	s.point.grade = "C"

	s.ViewPoint()

	// 값이 변경되지 않는다.
	// 포인터를 이해해야 한다.
	s.InputPoint("과학", "A")
	s.ViewPoint() 
}