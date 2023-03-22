package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var s Set
	s.Elements = make(map[string]struct{})
	return &s
}

func (s *Set) Add(elm string) {
	s.Elements[elm] = struct{}{}
}

func (s *Set) Delete(elm string) error {
	if _, exists := s.Elements[elm]; !exists {
		return errors.New("element not present in set")
	}

	delete(s.Elements, elm)
	return nil
}

func (s *Set) Contains(elm string) bool {
	_, exists := s.Elements[elm]
	return exists
}

func (s *Set) List() {
	fmt.Println("print set:")
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}
