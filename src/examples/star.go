package main

import "fmt"

func main() {
	star1() // OK
	fmt.Println()

	star2() // OK
	fmt.Println()

	star3() // OK
	fmt.Println()

	star4()
	fmt.Println()

	star5()
}


/*--------
*
* *
* * *
* * * *
* * * * *
--------*/
func star1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i + 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}


/*--------
* * * * * 
* * * *
* * *
* *
*
--------*/
func star2() {
	var lineNumber int = 5

	for i := 0; i < lineNumber; i++ {
		for j := 0; j < lineNumber - i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}


/*-----
*      
* *
* * *
* *
*
-----*/
func star3() {
	var lineNumber int = 5

	for i := 0; i < lineNumber; i++ {

		if i < 3 {

			for j := 0; j < i + 1; j++ {
				fmt.Print("*")
			}

			fmt.Println()
		} else {

			for k := 0; k < lineNumber - i; k++ {
				fmt.Print("*")
			}

			fmt.Println()
		}
	}
}


/*---------
    *      
  * * * 
* * * * *
----------*/
func star4() {
	var lineNumber int = 5

	for i := 0; i < lineNumber; i++ {

		for j := 0; j < (lineNumber - 1) - i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < i + (i + 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}


/*---------
    *      
  * * * 
* * * * *
	* * *
	  *
----------*/
func star5() {

}
// fmt.Println() : 입력받은 값을 출력하고 줄바꿈
// fmt.Print() : 입력받은 내용을 출력 ( 줄바꿈 x )