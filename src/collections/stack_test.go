package collections

import (
	"testing"
)

func Test_depth_maxdepth_after_init(t *testing.T) {
	s := Stack[int]{}
	//s.Push(0)
	if s.Depth() != 0 {
		t.Error("Stack initialized with incorrect depth. Expected 0, got", s.Depth())
	}
	if s.MaxDepth() != 0 {
		t.Error("Stack initialized with incorrect max depth. Expected 0, got", s.MaxDepth())
	}
}

func Test_push_pop(t *testing.T) {

	s := Stack[int]{}
	s.Push(0)

	if s.Depth() != 1 {
		t.Error("01: Stack incorrect depth. Expected 1, got", s.Depth())
	}
	if s.MaxDepth() != 1 {
		t.Error("01: Stack incorrect max depth. Expected 1, got", s.MaxDepth())
	}

	s.Push(1)

	if s.Depth() != 2 {
		t.Error("02: Stack incorrect depth. Expected 2, got", s.Depth())
	}
	if s.MaxDepth() != 2 {
		t.Error("02: Stack incorrect max depth. Expected 2, got", s.MaxDepth())
	}

	s.Pop()
	if s.Depth() != 1 {
		t.Error("03: Stack incorrect depth. Expected 2, got", s.Depth())
	}
	if s.MaxDepth() != 2 {
		t.Error("03: Stack incorrect max depth. Expected 2, got", s.MaxDepth())
	}

	s.Push(2)
	s.Push(3)

	s.Pop()
	if s.Depth() != 2 {
		t.Error("04: Stack incorrect depth. Expected 2, got", s.Depth())
	}
	if s.MaxDepth() != 3 {
		t.Error("04: Stack incorrect max depth. Expected 3, got", s.MaxDepth())
	}

}

func Test_peek_vs_pop(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v1 := s.Peek()
	if v1 != 3 {
		t.Error("Peek returns wrong value, expected 2, got", v1)
	}

	v2 := s.Pop()
	if v2 != v1 {
		t.Error("Pop returns other value than last peek, expected ", v1, " got ", v2)
	}
}

func Test_content_is_a_copy(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	c := s.Content()
	c[2] = 5
	if s.Peek() != 31 {
		t.Error("Content() returns a mutable slice")
	}
}
