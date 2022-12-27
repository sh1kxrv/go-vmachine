package cpu

import "errors"

type Stack struct {
	Entries []int
}

func NewStack() *Stack {
	return &Stack{
		Entries: []int{},
	}
}

func (s *Stack) Size() int {
	return len(s.Entries)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Entries) == 0
}

func (s *Stack) Push(entry int) {
	s.Entries = append(s.Entries, entry)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	entry := s.Entries[len(s.Entries)-1]
	s.Entries = s.Entries[:len(s.Entries)-1]
	return entry, nil
}
