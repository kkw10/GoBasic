package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Result struct {
	strikes int
	balls int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cnt := 0 // 시도 횟수

	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumbers()

	for {
		cnt++

		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S 인가?
		if IsGameEnd(result) {
			// 게임 끝
			break
		}		
	}

	// 게임 끝. 몇번만에 맞췄는지 출력한다.
	fmt.Printf("%d 번만에 맞췄습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0 ~ 9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3]int

	for i := 0; i < 3; i++ {

		for {
			n := rand.Intn(10)

			isDuplicated := false

			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					isDuplicated = true
					break
				}
			}
			
			if !isDuplicated {
				rst[i] = n
				break
			}
		}
	}

	fmt.Println(rst);
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0 ~ 9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")

		var numb int
		_, err := fmt.Scanf("%d\n", &numb)
	
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for numb > 0 {
			n := numb % 10
			numb = numb / 10

			isDuplicated := false

			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					isDuplicated = true
					break
				}
			}
			
			if isDuplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}

		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요")
			success = false
		}

		if !success {
			continue
		}

		break
	}

	rst[0], rst[2] = rst[2], rst[0]

	fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return Result{strikes, balls}
}

func IsGameEnd(result Result) bool {
	// 비교 결과가 3s인지 확인한다.
	return result.strikes == 3
}

func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}