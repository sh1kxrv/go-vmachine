package cpu

import "errors"

type Stack struct {
	Entries []int64
}

func NewStack() *Stack {
	return &Stack{
		Entries: []int64{},
	}
}

func (s *Stack) Size() int {
	return len(s.Entries)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Entries) == 0
}

func (s *Stack) Push(entry int64) {
	s.Entries = append(s.Entries, entry)
}

func (s *Stack) Pop() (int64, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	entry := s.Entries[len(s.Entries)-1]
	s.Entries = s.Entries[:len(s.Entries)-1]
	return entry, nil
}
