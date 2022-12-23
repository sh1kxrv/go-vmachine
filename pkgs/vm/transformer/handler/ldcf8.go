package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdcf8(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdcf8: %v", instr))
}
