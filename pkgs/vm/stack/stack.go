package stack

import "go-vmachine/pkgs/vm/instruction"

type Stack struct {
	Instructions []*instruction.Instruction
}

func NewStack() *Stack {
	return &Stack{
		Instructions: make([]*instruction.Instruction, 0),
	}
}

func (s *Stack) Push(instructions ...*instruction.Instruction) *Stack {
	s.Instructions = append(instructions, s.Instructions...)
	return s
}

func (s *Stack) Pop() *instruction.Instruction {
	if len(s.Instructions) == 0 {
		panic("stack is empty")
	}
	if len(s.Instructions) == 1 {
		popped := s.Instructions[0]
		s.Instructions = make([]*instruction.Instruction, 0)
		return popped
	}
	popped := s.Instructions[0]
	s.Instructions = s.Instructions[1:]
	return popped
}

func (s *Stack) Debug() {
	for _, instruction := range s.Instructions {
		println(instruction.String())
	}
}
