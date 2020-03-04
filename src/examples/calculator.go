package main

import (
	"fmt"
	"os" // 표준 입력을 받기 위함.
	"bufio" // 한줄씩 입력받기.
	"strings" // 입력받은 문자열에서 불순물 제거용.
	"strconv" // 문자열 숫자로 전환
)

func main() {
	fmt.Println("숫자를 입력하세요...")

	// 표준입력으로 입력을 받는다.
	reader := bufio.NewReader(os.Stdin)

	// 받은 입력에서 '\n' 까지 읽는다. 즉 한줄을 읽는다.
	// _는 일종의 이름없는 변수로 여기에 할당되는 값은 처리하지 않겠다는 의미이다. (여기서는 에러가 할당된다.)
	line, _ := reader.ReadString('\n')

	// 해당 줄에서 공백을 제거한다.
	line = strings.TrimSpace(line)

	// 입력받은 문자열을 숫자로 전환한다.
	n1, _ := strconv.Atoi(line)

	// 두번 쩨 숫자를 입력받는다.
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요...")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if (line == "+") {
		fmt.Printf("%d + %d = %d이다.", n1, n2, n1 + n2)
	} else if (line == "-") {
		fmt.Printf("%d - %d = %d이다.", n1, n2, n1 - n2)
	} else if (line == "*") {
		fmt.Printf("%d * %d = %d이다.", n1, n2, n1 * n2)
	} else if (line == "/") {
		fmt.Printf("%d / %d = %d이다.", n1, n2, n1 / n2)
	} else {
		fmt.Printf("유효하지 않은 연산자 입니다.")
	}
}