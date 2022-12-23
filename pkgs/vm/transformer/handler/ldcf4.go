package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdcf4(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdcf4: %v", instr))
}
