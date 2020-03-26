package main

import (
	"fmt"
	"data_structure"
)

func main() {
	tree := data_structure.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()

	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6")
		fmt.Println(cnt)
	} else {
		fmt.Println("not found 6")
		fmt.Println(cnt)
	}

	if found, cnt := tree.Search(12); found {
		fmt.Println("found 11")
		fmt.Println(cnt)
	} else {
		fmt.Println("not found 11")
		fmt.Println(cnt)
	}
}