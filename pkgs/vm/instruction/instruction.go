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

func NewNumber(value interface{}) *Instruction {
	switch value.(type) {
	case int:
		return NewInstruction(OpCodeLdci4, value)
	case int8:
		return NewInstruction(OpCodeLdci1, value)
	case int16:
		return NewInstruction(OpCodeLdci2, value)
	case int32:
		return NewInstruction(OpCodeLdci4, value)
	case int64:
		return NewInstruction(OpCodeLdci8, value)
	case float32:
		return NewInstruction(OpCodeLdcf4, value)
	case float64:
		return NewInstruction(OpCodeLdcf8, value)
	default:
		panic("invalid number type")
	}
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

func (i *Instruction) IsOperator() bool {
	return i.OpCode == OpCodeOperator
}

func (i *Instruction) IsNumber() bool {
	return i.OpCode >= OpCodeLdci1 && i.OpCode <= OpCodeLdcf8
}

func (i *Instruction) SetOperand(operand interface{}) {
	i.Operand = operand
}

func (i *Instruction) SetOpCode(opcode OpCode) {
	i.OpCode = opcode
}
