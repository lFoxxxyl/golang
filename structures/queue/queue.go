package main

type queue struct {
	items []int
	count int
}

func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
	q.count++
}

func (q *queue) dequeue() int {
	if q.count == 0 {
		panic("The stack is empty")
	}
	result := q.items[0]
	q.items = q.items[1:]
	q.count--
	return result
}
func (q *queue) peek() int {
	if q.count == 0 {
		panic("The stack is empty")
	}
	return q.items[0]
}
