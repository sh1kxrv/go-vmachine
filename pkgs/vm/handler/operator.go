package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

func OperatorHandler(stack *stack.Stack, instr *instruction.Instruction) {
	v1, v2 := stack.Pop(), stack.Pop()
	if v1 == nil || v2 == nil {
		panic("expected two values on stack")
	} else if v1.OpCode != instruction.OpCodeNumber || v2.OpCode != instruction.OpCodeNumber {
		panic("expected two numbers on stack")
	}

	stack.Pop() // pop self instruction

	var result *instruction.Instruction
	operator := instr.Operand.(instruction.OpCode)
	switch operator {
	case instruction.OpCodeAdd:
		result = instruction.NewNumber(v1.Operand.(int) + v2.Operand.(int))
	case instruction.OpCodeSub:
		result = instruction.NewNumber(v1.Operand.(int) - v2.Operand.(int))
	default:
		panic(fmt.Sprintf("invalid instruction %v", instr))
	}

	stack.Push(result)
}
