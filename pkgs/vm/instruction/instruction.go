package instruction

import "fmt"

type Instruction struct {
	Operand interface{}
	OpCode  OpCode
}

func NewInstruction(opcode OpCode, operand interface{}) *Instruction {
	return &Instruction{
		OpCode:  opcode,
		Operand: operand,
	}
}

func NewNumber(value int) *Instruction {
	return NewInstruction(OpCodeNumber, value)
}

func NewReturn() *Instruction {
	return NewInstruction(OpCodeRet, nil)
}

func NewOperator(operator OpCode) *Instruction {
	return NewInstruction(OpCodeOperator, operator)
}

func (i *Instruction) String() string {
	return fmt.Sprintf("OpCode: %v | Operand: %v", i.OpCode.String(), i.Operand)
}
