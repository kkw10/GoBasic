package main

import "fmt"

func main() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d단\n", dan)

		for i := 1; i <= 9; i++ {
			fmt.Printf("%d * %d = %d\n", dan, i, dan * i)
		}

		fmt.Println()
	}
}