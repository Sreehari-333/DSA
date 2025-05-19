package DsaToolKit

import "fmt"

type QueueNode struct {
	data int
	next *QueueNode
}

type queue struct {
	front *QueueNode
	rear  *QueueNode
}

func (q *queue) enQueue(val int) {
	newNode := &QueueNode{data: val}

	if q.front == nil {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *queue) deQueue() int {

	if q.front == nil {
		fmt.Println("Queue is empty")
		return 0
	}

	val := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return val
}

func (q *queue) display() {

	if q.rear == nil {
		fmt.Println("Queue is empty")
	}

	current := q.front

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func Queue() {

	q1 := &queue{}

	q1.enQueue(10)
	q1.enQueue(20)

	fmt.Println("After inserting")
	q1.display()

	fmt.Println("After DeQueue")
	fmt.Println("Removed element", q1.deQueue())
	q1.display()

}
