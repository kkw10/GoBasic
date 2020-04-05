package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["abc"] = "bbbb"
	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])

	// 설정되어 있지 않은 값을 조회하면 해당 자료형의 기본 값이 출력된다.
	// 이러한 특성을 통해서 map에 값이 있는지 없는지 확인할 수 있다.
	m2 := make(map[int]bool)
	m2[4] = true
	fmt.Println(m2[6], m2[4])

	// 하지만 기본형을 할당한 경우 앞선 방식으로는 확인할 수 없다.
	// 대신 map에서 제공하는 2번째 값을 통해서 확인할 수 있다.
	v, ok := m1[53]
	v1, ok1 := m1[1]
	fmt.Println(v, ok, v1, ok1)

	delete(m1, 53)
	v, ok = m1[53]
	fmt.Println(v, ok)

	m3 := make(map[int]int)
	m3[2] = 123
	m3[10] = 456
	m3[4] = 789

	for key, value := range m3 {
		// map의 값은 정렬되어 있지 않기 때문에 순서없이 무작위로 반환된다.
		// 정렬된 값을 얻고 싶다면 key값을 갖고 따로 정렬을 해야한다.
		fmt.Println(key, " = ", value)
	}
}