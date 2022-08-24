package main

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head  *node
	tail  *node
	count int
}

func (l *linkedList) add(i int) {
	n := node{value: i}
	if l.head == nil {
		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}
	l.count++
}

func (l *linkedList) remove(i int) bool {
	var previous *node = nil
	var current *node = l.head
	for current != nil {
		if current.value == i {
			if previous != nil {
				previous.next = current.next
				if current.next == nil {
					l.tail = previous
				}
			} else {
				l.head = l.head.next
				if l.head == nil {
					l.tail = nil
				}
			}
			l.count--

			return true
		}
		previous = current
		current = current.next
	}
	return false
}

func (l linkedList) contains(i int) bool {
	current := l.head
	for current != nil {
		if current.value == i {
			return true
		}
		current = current.next
	}
	return false
}

func (l *linkedList) clear() {
	l.head = nil
	l.tail = nil
	l.count = 0
}

func (l linkedList) copyTo(a *[]int) {
	current := l.head
	for current != nil {
		*a = append(*a, current.value)
		current = current.next
	}
}
