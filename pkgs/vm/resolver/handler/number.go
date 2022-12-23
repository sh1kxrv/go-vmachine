package handler

import (
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

func NumberHandler(s *stack.Stack, instr *instruction.Instruction) {
	s.Push(
		stack.NewStackValue(instr.Operand, instr.OpCode),
	)
}
