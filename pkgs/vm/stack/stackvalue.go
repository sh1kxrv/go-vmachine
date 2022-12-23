package stack

import "go-vmachine/pkgs/vm/instruction"

type StackValue struct {
	Value  interface{}
	OpCode instruction.OpCode
}

func NewStackValue(value interface{}, opCode instruction.OpCode) *StackValue {
	return &StackValue{
		Value:  value,
		OpCode: opCode,
	}
}

func (s *StackValue) IsNumber() bool {
	return s.OpCode >= instruction.OpCodeLdci1 && s.OpCode <= instruction.OpCodeLdcf8
}
