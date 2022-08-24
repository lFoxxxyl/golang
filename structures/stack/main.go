package main

import "fmt"

func main(){
	var s stack
	fmt.Println(s.count, s.items)

	s.push(1)
	fmt.Println(s.count, s.items)

	s.pop()
	fmt.Println(s.count, s.items)

	s.push(1)
	fmt.Println(s.count, s.items)
	s.push(2)
	fmt.Println(s.count, s.items)

	s.peek()
	fmt.Println(s.count, s.items)
}