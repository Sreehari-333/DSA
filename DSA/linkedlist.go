package DsaToolKit

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
}

func (l *Linkedlist) insert(val int) {

	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (l *Linkedlist) delete(value int) {

	if l.head == nil {
		fmt.Println("List is empty")
	}

	if l.head.data == value {
		l.head = l.head.next
	}

	prev := l.head
	current := l.head.next

	for current != nil {
		if current.data == value {
			prev.next = current.next
		}
		prev = current
		current = current.next
	}

}

func (l *Linkedlist) toSlice() []int {

	var result = []int{}

	if l.head == nil {
		fmt.Println("List is empty")
	}

	current := l.head

	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (l *Linkedlist) display() {

	if l.head == nil {
		fmt.Println("List is empty")
	}

	current := l.head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func LinkedList() {

	l1 := &Linkedlist{}

	l1.insert(10)
	l1.insert(20)
	l1.insert(30)
	fmt.Println("After inserting")
	l1.display()

	fmt.Println("Converting list to slice")
	fmt.Println(l1.toSlice())

	fmt.Println("After delete a element")
	l1.delete(30)
	l1.display()

}
