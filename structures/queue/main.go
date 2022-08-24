package main

import "fmt"

func main(){
	var q queue
	fmt.Println(q.count, q.items)

	q.enqueue(1)
	fmt.Println(q.count, q.items)

	q.dequeue()
	fmt.Println(q.count, q.items)

	q.enqueue(1)
	fmt.Println(q.count, q.items)
	q.enqueue(2)
	fmt.Println(q.count, q.items)

	q.peek()
	fmt.Println(q.count, q.items)
}