package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdcf4(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdcf4: %v", instr))
}
