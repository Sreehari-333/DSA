package DsaToolKit

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

type stack struct {
	top *StackNode
}

func (s *stack) push(val int) {

	newNode := &StackNode{data: val, next: s.top}
	s.top = newNode
}

func (s *stack) pop() int {

	if s.top == nil {
		fmt.Println("Stack is empty")
	}

	val := s.top.data
	s.top = s.top.next

	return val
}

func (s *stack) peek() {
	if s.top == nil {
		fmt.Println("Stack is empty")
	}

	fmt.Println(s.top.data)
}

func (s *stack) display() {

	if s.top == nil {
		fmt.Println("Stack is empty")
	}

	current := s.top

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func Stack() {

	s1 := &stack{}

	s1.push(10)
	s1.push(20)
	s1.push(30)
	fmt.Println("After inserting")
	s1.display()

	fmt.Println("Removing the last element")
	fmt.Println("Removed element : ", s1.pop())
	s1.display()

	fmt.Println("Viewing the last element")
	s1.peek()

}
