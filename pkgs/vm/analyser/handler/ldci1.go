package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdci1(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdci1: %v", instr))
}
