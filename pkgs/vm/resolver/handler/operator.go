package handler

import (
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/resolver/handler/operator"
	"go-vmachine/pkgs/vm/stack"
)

func OperatorHandler(stack *stack.Stack, instr *instruction.Instruction) {
	v1, v2 := stack.Pop(), stack.Pop()
	if v1 == nil || v2 == nil {
		panic("expected two values on stack")
	} else if !v1.IsNumber() || !v2.IsNumber() {
		panic("expected two numbers on stack")
	}

	operator_by_operand := instr.Operand.(instruction.OpCode)
	if operator_by_operand == instruction.OpCodeAdd {
		operator.OperatorAdd(stack, v1, v2)
	} else if operator_by_operand == instruction.OpCodeSub {
		operator.OperatorSub(stack, v1, v2)
	} else if operator_by_operand == instruction.OpCodeMul {
		operator.OperatorMul(stack, v1, v2)
	} else if operator_by_operand == instruction.OpCodeDiv {
		operator.OperatorDiv(stack, v1, v2)
	}
}
