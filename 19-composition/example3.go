package main

type StringStack struct {
	data []string // data is private, does not expose the implementation of StringStack.
}

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
	// Avoids exposing that there is a slice.
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}

	panic("Pop from empty stack.") // Control the message: no mention of a slice.
}

type IntList struct {
	Value int
	Tail  *IntList
}

// Make nil useful.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}
