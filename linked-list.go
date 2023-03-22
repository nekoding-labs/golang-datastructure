package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Size int
}

type Node struct {
	Value string
	Next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(item string) {
	node := Node{
		Value: item,
		Next:  l.Head,
	}

	l.Head = &node
	l.Size++
}

func (l *LinkedList) RemoveFirst() (bool, error) {
	if l.Size == 0 {
		return false, errors.New("linked list is empty")
	}

	l.Head = l.Head.Next
	l.Size--

	return true, nil
}

func (l LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v \n", current)
		current = current.Next
	}
}

func (l LinkedList) Search(s string) (*Node, error) {
	current := l.Head
	for current != nil {
		if current.Value == s {
			return current, nil
		}
		current = current.Next
	}

	return nil, errors.New("linked list is empty")
}

func (l *LinkedList) RemoveItem(s string) (bool, error) {
	previous := l.Head
	current := l.Head

	for current != nil {
		if current.Value == s {
			previous.Next = current.Next
			l.Size--
			return true, nil
		}

		previous = current
		current = current.Next
	}

	return false, errors.New("linked list is empty")
}
