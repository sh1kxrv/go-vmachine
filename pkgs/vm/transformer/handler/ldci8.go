package handler

import (
	"go-vmachine/pkgs/logger"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci8(instr *instruction.Instruction)  {
	v1 := instr.Operand.(int64)
	if v1 >= -127 && v1 <= 128 {
		instr.SetOpCode(instruction.OpCodeLdci1)
		instr.SetOperand(int8(v1))
		logger.Warn("Transformed -> Ldci1", v1)
	} else if v1 >= -32768 && v1 <= 32767 {
		instr.SetOpCode(instruction.OpCodeLdci2)
		instr.SetOperand(int16(v1))
		logger.Warn("Transformed -> Ldci2", v1)
	} else if v1 >= -2147483648 && v1 <= 2147483647 {
		instr.SetOpCode(instruction.OpCodeLdci4)
		instr.SetOperand(int(v1))
		logger.Warn("Transformed -> Ldci4", v1)
	}
}
