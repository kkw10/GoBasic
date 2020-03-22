package data_structure

type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	return &Stack{ll: &LinkedList{}}
}

func (s *Stack) IsEmpty() bool {
	return s.ll.IsEmpty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	lastValue := s.ll.ReturnLastValue()
	s.ll.PopLastValue()
	return lastValue
}