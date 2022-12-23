package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformLdci2(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformLdci2: %v", instr))
}
