package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci4(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdci4: %v", instr))
}
