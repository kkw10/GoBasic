package data_structure

import "fmt"

// 패키지에서 소문자로 시작되는 것은 외부로 공개되지 않는다.
// 반대로 대문자로 시작된는 것은 외부에서 사용할 수 있다.

type Node struct {
	Next *Node
	Prev *Node
	Val int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) IsEmpty() bool {
	return l.Root == nil
} 

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next

		if l.Root != nil {
			l.Root.Prev = nil
		}

		node.Next = nil
		return
	}

	prev := node.Prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
	} else {
		node.Prev = nil
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
	}

	node.Next = nil
}

func (l *LinkedList) ReturnLastValue() int {
	if l.Tail != nil { return l.Tail.Val }
	return 0
}
func (l *LinkedList) ReturnFirstValue() int {
	if l.Root != nil { return l.Root.Val }
	return 0
}

func (l *LinkedList) PopLastValue() {
	if l.Tail == nil { return }
	l.RemoveNode(l.Tail)
}
func (l *LinkedList) PopFirstValue() {
	if l.Root == nil { return }
	l.RemoveNode(l.Root)
}

func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}

	fmt.Printf("%d\n", node.Val)
}
func (l *LinkedList) PrintRevers() {
	node := l.Tail

	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}

	fmt.Printf("%d\n", node.Val)
}