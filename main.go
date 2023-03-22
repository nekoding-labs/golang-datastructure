package main

import "fmt"

func main() {

	// Set
	set := NewSet()
	set.Add("enggar")
	set.Add("rachmad")
	set.Add("tivandi")

	set.List()
	set.Add("enggar")
	set.List()

	set.Delete("tivandi")
	set.List()

	// Queue
	q := NewQueue()
	q.Enqueue(1, false)
	q.Enqueue(2, true)
	q.Enqueue(3, true)

	fmt.Println("before dequeue:", q)
	q.Dequeue()
	fmt.Println("after dequeue:", q)

	fd, _ := q.Peek()
	fmt.Println("peek queue :", fd)

	// Stack
	s := NewStack()
	s.Push(1)
	s.Push(2)
	fmt.Println(s)

	s.Pop()
	fmt.Println(s)
	fmt.Println(s.IsEmpty())

	st, _ := s.Peek()
	fmt.Println(st)
	fmt.Printf("Stack length is %d\n", s.Length())

	// Linked list
	ll := NewLinkedList()
	ll.Add("hello")
	ll.Add("world")
	ll.RemoveFirst()

	ll.Add("enggar")
	ll.Add("tivandi")

	found, err := ll.Search("enggar")
	if err == nil {
		fmt.Printf("search data found, %v \n", found)
	}

	_, err = ll.RemoveItem("enggar")
	if err == nil {
		fmt.Println("remove item success")
	}

	ll.List()
}
