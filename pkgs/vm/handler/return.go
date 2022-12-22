package handler

import (
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

func ReturnHandler(stack *stack.Stack, instruction *instruction.Instruction) {
	v := stack.Pop()
	if v == nil {
		panic("invalid instructions")
	}
	instruction.Operand = v
	// ...
}
