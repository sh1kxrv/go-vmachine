package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci1(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdci1: %v", instr))
}
