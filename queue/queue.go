package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue remove a value at the front
func (q *Queue) Dequeue() {
	q.items = q.items[1:]
}

func main() {
	myQueue := &Queue{}
	myQueue.Enqueue(25)
	myQueue.Enqueue(45)
	myQueue.Enqueue(23)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
