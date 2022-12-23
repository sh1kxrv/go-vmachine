package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci8(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdci8: %v", instr))
}
