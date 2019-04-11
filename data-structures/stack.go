package main

import (
	"errors"
	"fmt"
)

type Container interface {
	Size() int64
	Empty() bool
	Clear()
}

type Stack interface {
	Container
	Push(e interface{})
	Pop() (interface{}, error)
	Top() (interface{}, error)
}

type ArrayStack struct {
	store []interface{}
}

func (s *ArrayStack) Size() int {
	return len(s.store)
}
func (s *ArrayStack) Empty() bool {
	return len(s.store) == 0
}
func (s *ArrayStack) Clear() {
	s.store = make([]interface{}, 0, 10)
}

func (s *ArrayStack) Push(e interface{}) {
	s.store = append(s.store, e)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Pop: the stack cannot be empty")
	}
	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return result, nil
}

func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Top: stack cannot be empty")
	}
	return s.store[len(s.store)-1], nil
}

func main() {
	var as ArrayStack

	as.Push("foo")
	as.Push("bar")
	as.Pop()

	// fmt.Println(as.Size())
	fmt.Println(as.Top())
}
