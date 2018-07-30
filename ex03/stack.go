package stack

type (
	Stack struct {
		head *node
		size int
	}

	node struct {
		prev  *node
		value interface{}
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Push(v interface{}) {
	n := &node{this.head, v}
	this.head = n
	this.size++
}

func (this *Stack) Pop() interface{} {
	if this.size == 0 {
		return nil
	}

	n := this.head
	this.head = n.prev
	this.size--
	return n.value
}
