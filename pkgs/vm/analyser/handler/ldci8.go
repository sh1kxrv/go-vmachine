package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdci8(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdci8: %v", instr))
}
