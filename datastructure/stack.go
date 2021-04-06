package datastructure

// Stack represent stack that hold slice
type Stack struct {
	items []int
}

// Push to push slice into stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop to pop slice out of stack
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}
