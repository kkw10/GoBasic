package main

import (
	"fmt"
	"data_structure"
)

func main() {
	fmt.Println("abcde =>", data_structure.RollingHash("abcde"))
	fmt.Println("abcde =>", data_structure.RollingHash("abcde"))
	fmt.Println("abcdf =>", data_structure.RollingHash("abcdf"))
	fmt.Println("tbcde =>", data_structure.RollingHash("tbcde"))

	fmt.Println("/////////////////////")

	m := data_structure.CreateMap()
	m.Add("AAA", "01011111111")
	m.Add("BBB", "01022222222")
	m.Add("CCC", "01033333333")

	fmt.Println("AAA =>", m.Get("AAA"))
	fmt.Println("BBB =>", m.Get("BBB"))
	fmt.Println("CCC =>", m.Get("CCC"))
	fmt.Println("DDD =>", m.Get("DDD"))
}