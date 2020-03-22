package data_structure

type Queue struct {
	ll *LinkedList
}

func (q *Queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}

func NewQueue() *Queue {
	return &Queue{ ll: &LinkedList{} }
}

func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() int {
	firstValue := q.ll.ReturnFirstValue()
	q.ll.PopFirstValue()
	return firstValue
}