package operator

import (
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

func OperatorDivLdci1(v1, v2 int8) int8 {
	return v1 / v2
}

func OperatorDivLdci2(v1, v2 int16) int16 {
	return v1 / v2
}

func OperatorDivLdci4(v1, v2 int) int {
	return v1 / v2
}

func OperatorDivLdci8(v1, v2 int64) int64 {
	return v1 / v2
}

func OperatorDivLdcf4(v1, v2 float32) float32 {
	return v1 / v2
}

func OperatorDivLdcf8(v1, v2 float64) float64 {
	return v1 / v2
}

func OperatorDiv(s *stack.Stack, v1, v2 *stack.StackValue) {
	if v1.OpCode != v2.OpCode {
		panic("expected type cast to be the same")
	}
	switch v1.OpCode {
	case instruction.OpCodeLdci1:
		s.Push(stack.NewStackValue(OperatorDivLdci1(v1.Value.(int8), v2.Value.(int8)), instruction.OpCodeLdci1))
	case instruction.OpCodeLdci2:
		s.Push(stack.NewStackValue(OperatorDivLdci2(v1.Value.(int16), v2.Value.(int16)), instruction.OpCodeLdci2))
	case instruction.OpCodeLdci4:
		s.Push(stack.NewStackValue(OperatorDivLdci4(v1.Value.(int), v2.Value.(int)), instruction.OpCodeLdci4))
	case instruction.OpCodeLdci8:
		s.Push(stack.NewStackValue(OperatorDivLdci8(v1.Value.(int64), v2.Value.(int64)), instruction.OpCodeLdci8))
	case instruction.OpCodeLdcf4:
		s.Push(stack.NewStackValue(OperatorDivLdcf4(v1.Value.(float32), v2.Value.(float32)), instruction.OpCodeLdcf4))
	case instruction.OpCodeLdcf8:
		s.Push(stack.NewStackValue(OperatorDivLdcf8(v1.Value.(float64), v2.Value.(float64)), instruction.OpCodeLdcf8))
	default:
		panic("expected two numbers on stack")
	}
}
