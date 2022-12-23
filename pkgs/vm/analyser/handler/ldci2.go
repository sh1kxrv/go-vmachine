package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdci2(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdci2: %v", instr))
}
