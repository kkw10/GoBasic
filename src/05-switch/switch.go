package main

import "fmt"

func main() {
	a := 33

	switch (a) {
	case 31:
		fmt.Println("a = 31")
	case 32:
		fmt.Println("a = 32")
	case 33:
		fmt.Println("a = 33")
	}

	switch {
	case a > 40:
		fmt.Println("a는 40보다 크다.")
	case a < 20:
		fmt.Println("a는 20보다 작다.")
	case a > 30:
		fmt.Println("a는 30보다 크다.")
	case a == 33: // 앞서 걸린 케이스만 실행된다.
		fmt.Println("a == 33")
	}
}