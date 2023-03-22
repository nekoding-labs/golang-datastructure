package main

import "errors"

type Stack struct {
	Elements []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item int) {
	s.Elements = append(s.Elements, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastElemIndex := len(s.Elements) - 1
	lastElemIndex, s.Elements = s.Elements[lastElemIndex], s.Elements[:lastElemIndex]

	return lastElemIndex, nil
}

func (s Stack) IsEmpty() bool {
	return len(s.Elements) < 1
}

func (s Stack) Peek() (int, error) {
	if len(s.Elements) < 1 {
		return 0, errors.New("stack is empty")
	}

	return s.Elements[len(s.Elements)-1], nil
}

func (s Stack) Length() int {
	return len(s.Elements)
}
