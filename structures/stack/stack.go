package main

type stack struct{
	items []int
	count int
}
func(s *stack)push(i int){
	s.items = append(s.items, i)
	s.count+=1
}

func(s *stack)pop() int{
	if s.count==0{
		panic("The stack is empty")
	}
	result:= s.items[s.count-1]
	s.items=s.items[:len(s.items)-1]
	s.count--
	return result
}

func(s *stack)peek()int{
	if s.count==0{
		panic("The stack is empty")
	}

	return s.items[s.count-1]
}