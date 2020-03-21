package main

import "fmt"

type Node struct {
	next *Node
	val int
}

func AddNode(root *Node, val int) {
	var tail *Node
	tail = root

	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{ val: val }
	tail.next = node
}

// Go로 Linked List 구현해보기 way1
func main() {
	var root *Node
	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}