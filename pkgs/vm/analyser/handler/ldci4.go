package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdci4(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdci4: %v", instr))
}
