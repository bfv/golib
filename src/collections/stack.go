package collections

import (
	"fmt"
)

// generic Stack implmentation providing stack-like Pop/Push function
// the fields are only accessible via the functions
type Stack[T any] struct {
	values   []T
	maxDepth int
}

// Push adds a value on top of the stack and updates maxDepth (if applicable)
func (s *Stack[T]) Push(val T) error {
	s.values = append(s.values, val)
	if len(s.values) > s.maxDepth {
		s.maxDepth = len(s.values)
	}
	return nil
}

// Pop removed and returns the top value from the Stack.
func (s *Stack[T]) Pop() T {
	v := s.values[len(s.values)-1:]
	s.values = s.values[:len(s.values)-1]
	return v[0]
}

// Peek returns the last inserted value of the stack (no removal)
func (s *Stack[T]) Peek() interface{} {
	return s.values[len(s.values)-1:][0]
}

// returns the current depth of the stack
func (s *Stack[T]) Depth() int {
	return len(s.values)
}

// MaxDepth return the maximum amount of values this Stack reached
func (s *Stack[T]) MaxDepth() int {
	return s.maxDepth
}

// Returns true is the stack holds no values
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// String output the Stack contents prefixed with an index
func (s *Stack[T]) String() string {
	var str string
	for i, v := range s.values {
		str = fmt.Sprintf("%2d: %v\n%s", i, v, str)
	}
	return str
}

// Content returns a copy of the slice holding the values
func (s *Stack[T]) Content() []T {
	c := make([]T, len(s.values))
	copy(c, s.values)
	return c
}
