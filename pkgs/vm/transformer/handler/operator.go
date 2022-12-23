package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func TransformOperator(instr *instruction.Instruction)  {
	println(fmt.Sprintf("TransformOperator: %v", instr))
}
