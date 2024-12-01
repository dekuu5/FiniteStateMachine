package nfa

// Stack represents a stack data structure
type Stack struct {
	data []interface{}
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{data: []interface{}{}}
}

// Push adds an element to the stack
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

// Peek returns the top element without removing it
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
