package stack

import "fmt"

type Stack struct {
	StackData []*StackValue
}

func NewStack() *Stack {
	return &Stack{
		StackData: make([]*StackValue, 0),
	}
}

func (s *Stack) Push(data ...*StackValue) *Stack {
	s.StackData = append(s.StackData, data...)
	return s
}

func (s *Stack) Pop() *StackValue {
	if len(s.StackData) == 0 {
		return nil
	}
	if len(s.StackData) == 1 {
		popped := s.StackData[0]
		s.StackData = make([]*StackValue, 0)
		return popped
	}
	popped := s.StackData[0]
	s.StackData = s.StackData[1:]
	return popped
}

func (s *Stack) Debug() {
	for _, v := range s.StackData {
		println(fmt.Sprintf("%v", v))
	}
}
