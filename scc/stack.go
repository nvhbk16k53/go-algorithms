package main

import (
	"errors"
	"sync"
)

// Stack ...
type Stack struct {
	mutex sync.Mutex
	s     []interface{}
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{
		mutex: sync.Mutex{},
		s:     make([]interface{}, 0),
	}
}

func (s *Stack) empty() bool {
	return len(s.s) == 0
}

// Empty ...
func (s *Stack) Empty() bool {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	return s.empty()
}

func (s *Stack) push(v interface{}) {
	s.s = append(s.s, v)
}

// Push ...
func (s *Stack) Push(v interface{}) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	s.push(v)
}

func (s *Stack) pop() (interface{}, error) {
	if s.empty() {
		return nil, errors.New("stack is empty")
	}

	n := len(s.s)
	v := s.s[n-1]
	s.s = s.s[:n-1]
	return v, nil
}

// Pop ...
func (s *Stack) Pop() (interface{}, error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	return s.pop()
}
