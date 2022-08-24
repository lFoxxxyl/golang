package main

import "fmt"

func main() {

	l := linkedList{}

	l.add(1)
	l.add(2)

	var a []int
	l.copyTo(&a)
	fmt.Println(a)

	l.clear()
	fmt.Println(l.head, l.tail, l.count)

	l.add(1)
	l.add(2)
	if l.contains(1) {
		l.remove(1)
	}
	l.remove(2)
	fmt.Println(l.head, l.tail, l.count)
}
