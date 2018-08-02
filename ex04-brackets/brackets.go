//package main
package brackets

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

func (this *Stack) Len() int {
	return this.size
}

func Bracket(s string) (bool, error) {
	var stack *Stack = New()

	for _, val := range s {
		if val == '{' {
			stack.Push('}')
		} else if val == '(' {
			stack.Push(')')
		} else if val == '[' {
			stack.Push(']')
		} else {
			br := stack.Pop()
			if br == nil || br != val {
				return false, nil
			}
		}
	}
	return stack.Len() == 0, nil
}
