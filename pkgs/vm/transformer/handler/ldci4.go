package handler

import (
	"go-vmachine/pkgs/logger"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci4(instr *instruction.Instruction)  {
	v1 := instr.Operand.(int)
	if v1 >= -127 && v1 <= 128 {
		instr.SetOpCode(instruction.OpCodeLdci1)
		instr.SetOperand(int8(v1))
		logger.Warn("Transformed -> Ldci1", v1)
	} else if v1 >= -32768 && v1 <= 32767 {
		instr.SetOpCode(instruction.OpCodeLdci2)
		instr.SetOperand(int16(v1))
		logger.Warn("Transformed -> Ldci2", v1)
	}
}
