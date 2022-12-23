package handler

import (
	"go-vmachine/pkgs/logger"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci2(instr *instruction.Instruction)  {
	v1 := instr.Operand.(int16)
	if v1 >= -127 && v1 <= 128 {
		instr.SetOpCode(instruction.OpCodeLdci1)
		instr.SetOperand(int8(v1))
		logger.Warn("Transformed -> Ldci1", v1)
	}
}
