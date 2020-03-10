package main

import "fmt"

func main() {
	way1()
	way2()
}

func way1() {
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		temp[i] = arr[(len(arr) - 1) - i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}

	fmt.Println(arr)
}

func way2() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[(len(arr) - 1) - i] = arr[(len(arr) - 1) - i], arr[i]
	}

	fmt.Println(arr)
}